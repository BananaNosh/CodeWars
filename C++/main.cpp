#include <string>
#include <vector>
#include <memory>
#include <iostream>

using namespace std;

class Edge;

struct State {
    string name;
    vector<shared_ptr<Edge>> transitions;
};

struct Edge {
    string event;
    shared_ptr<State> destination;
};


shared_ptr<State> createAutomaton() {
    const shared_ptr<State> &CLOSED = make_shared<State>(State{"CLOSED"});
    const shared_ptr<State> &LISTEN = make_shared<State>(State{"LISTEN"});
    const shared_ptr<State> &SYN_SENT = make_shared<State>(State{"SYN_SENT"});
    const shared_ptr<State> &SYN_RCVD = make_shared<State>(State{"SYN_RCVD"});
    const shared_ptr<State> &ESTABLISHED = make_shared<State>(State{"ESTABLISHED"});
    const shared_ptr<State> &CLOSE_WAIT = make_shared<State>(State{"CLOSE_WAIT"});
    const shared_ptr<State> &LAST_ACK = make_shared<State>(State{"LAST_ACK"});
    const shared_ptr<State> &FIN_WAIT_1 = make_shared<State>(State{"FIN_WAIT_1"});
    const shared_ptr<State> &FIN_WAIT_2 = make_shared<State>(State{"FIN_WAIT_2"});
    const shared_ptr<State> &CLOSING = make_shared<State>(State{"CLOSING"});
    const shared_ptr<State> &TIME_WAIT = make_shared<State>(State{"TIME_WAIT"});

    CLOSED->transitions.push_back(make_shared<Edge>(Edge{"APP_PASSIVE_OPEN", LISTEN}));
    CLOSED->transitions.push_back(make_shared<Edge>(Edge{"APP_ACTIVE_OPEN", SYN_SENT}));
    LISTEN->transitions.push_back(make_shared<Edge>(Edge{"RCV_SYN", SYN_RCVD}));
    LISTEN->transitions.push_back(make_shared<Edge>(Edge{"APP_SEND", SYN_SENT}));
    LISTEN->transitions.push_back(make_shared<Edge>(Edge{"APP_CLOSE", CLOSED}));
    SYN_RCVD->transitions.push_back(make_shared<Edge>(Edge{"APP_CLOSE", FIN_WAIT_1}));
    SYN_RCVD->transitions.push_back(make_shared<Edge>(Edge{"RCV_ACK", ESTABLISHED}));
    SYN_SENT->transitions.push_back(make_shared<Edge>(Edge{"RCV_SYN", SYN_RCVD}));
    SYN_SENT->transitions.push_back(make_shared<Edge>(Edge{"RCV_SYN_ACK", ESTABLISHED}));
    SYN_SENT->transitions.push_back(make_shared<Edge>(Edge{"APP_CLOSE", CLOSED}));
    ESTABLISHED->transitions.push_back(make_shared<Edge>(Edge{"APP_CLOSE", FIN_WAIT_1}));
    ESTABLISHED->transitions.push_back(make_shared<Edge>(Edge{"RCV_FIN", CLOSE_WAIT}));
    FIN_WAIT_1->transitions.push_back(make_shared<Edge>(Edge{"RCV_FIN", CLOSING}));
    FIN_WAIT_1->transitions.push_back(make_shared<Edge>(Edge{"RCV_FIN_ACK", TIME_WAIT}));
    FIN_WAIT_1->transitions.push_back(make_shared<Edge>(Edge{"RCV_ACK", FIN_WAIT_2}));
    CLOSING->transitions.push_back(make_shared<Edge>(Edge{"RCV_ACK", TIME_WAIT}));
    FIN_WAIT_2->transitions.push_back(make_shared<Edge>(Edge{"RCV_FIN", TIME_WAIT}));
    TIME_WAIT->transitions.push_back(make_shared<Edge>(Edge{"APP_TIMEOUT", CLOSED}));
    CLOSE_WAIT->transitions.push_back(make_shared<Edge>(Edge{"APP_CLOSE", LAST_ACK}));
    LAST_ACK->transitions.push_back(make_shared<Edge>(Edge{"RCV_ACK", CLOSED}));
    return CLOSED;
}


std::string traverse_TCP_states(const std::vector<std::string> &events) {
    shared_ptr<State> state = createAutomaton();
    for (const auto& eventName: events) {
        bool found = false;
        for (const auto & transition : state->transitions) {
            if (transition->event == eventName) {
                state = transition->destination;
                found = true;
                break;
            }
        }
        if (!found) {
            return "ERROR";
        }
    }
    return state->name;
}


int main() {
    using vs = std::vector<std::string>;
    vs test1 = {"APP_ACTIVE_OPEN","RCV_SYN_ACK","RCV_FIN"};
    vs test2 = {"APP_PASSIVE_OPEN",  "RCV_SYN","RCV_ACK"};
    vs test3 = {"APP_ACTIVE_OPEN","RCV_SYN_ACK","RCV_FIN","APP_CLOSE"};
    vs test4 = {"APP_ACTIVE_OPEN"};
    vs test5 = {"APP_PASSIVE_OPEN","RCV_SYN","RCV_ACK","APP_CLOSE","APP_SEND"};
    std::cout << traverse_TCP_states(test1) << std::endl;
    std::cout << traverse_TCP_states(test2) << std::endl;
    std::cout << traverse_TCP_states(test3) << std::endl;
    std::cout << traverse_TCP_states(test4) << std::endl;
    std::cout << traverse_TCP_states(test5) << std::endl;

    return 0;
}
