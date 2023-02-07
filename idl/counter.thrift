namespace go douyin.rpc

service CounterService {
  CounterNopResponse Increment (1: CounterIncRequest Req);
  CounterGetResponse Fetch (1: CounterGetRequest Req);
}

struct CounterNopResponse {}
struct Increment {
  1: i64 Id;
  2: i8 Kind;
  3: i16 Delta;
}
struct CounterIncRequest {
  1: list<Increment> Actions;
}

struct CounterGetRequest {
  1: list<i64> Id;
  2: optional list<i8> Kinds;
}
struct KindCount {
  1: i8 Kind;
  2: i32 Count;
}
struct Counts {
  1: i64 Id;
  2: list<KindCount> KindCounts;
}
struct CounterGetResponse {
  1: list<Counts> Counters;
}