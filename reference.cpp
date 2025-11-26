/* ------------------- number theory ------------------- */ 
const int MOD = 1e9 + 7;
const int N = 1e7 + 5;

// spf - this is better,
// use this instead of plain sieve of eratosthenes
// once we precomputed spf, we can fatorise in O(logn) complexity, as compared to online sqrt(n)
// harmonic series nlogn most basic instead of starting from i*i looks good enough

vector<int> spf(N, 1);
for (int i = 2; i < N; i++) {
    if (spf[i] == 1) {
        spf[i] = i;
        for (int j = i * i; j < N; j += i) {
            spf[j] = i;
        }
    }
}

// binpow 
int binpow(int x, int y) {
    int res = 1;
    x = x % MOD; 
    while (y > 0) {  
        if (y & 1) 
            res = (1LL * res * x) % MOD;  
        y = y >> 1;
        x = (1LL * x * x) % MOD;   
    } 
    return res; 
}

// add, mul, sub, inv in MOD 
int add (int x, int y) {
    return (x + y) % MOD;
}
int mul (int x, int y ) {
    return (1LL * x * y) % MOD;
}
int sub (int x, int y) {
    return (x - y + MOD) % MOD;
}

int inv (int x) {
    return binpow(x, MOD-2);
}


// mobius inversion

/* ------------------- graph ------------------- */ 

// dijkstra
vector<int> d(n);
set<pair<int, int>> s;
s.insert({0, 0});
while (s.size()) {
    int u = s.begin()->second, du = s.begin()->first; s.erase(s.begin());
    for (auto v:adj[u]) {
        int dnew = du + calcDist(v, u);
        if (dnew < d[v]) {
            s.erase({d[v], v});
            d[v] = dnew;
            s.insert({d[v], v});
        } 
    }
}

/* ------------------- two pointers ------------------- */ 
// same for good inside and good outside
int lengthOfLongestSubstring(string s) {
    int n = s.size();
    int j = 0;
    map<char, int> m;
    auto good = [&] () {
        bool ok = 1;
        for (auto p:m) {
            ok &= (p.second<=1);
        }
        return ok;
    };
    int ans = 0;
    for (int i=0; i<n; i++){
        m[s[i]]++;
        while (j<=i && !good()) {
            m[s[j]]--; j++;
        }
        ans = max(ans, i-j+1);
    }
    return ans;
}

// ternary search 

int l = 0, r = n-1; // both inclusive of the range  
while (r - l > 2) {
    int x = l + (r - l) / 3;
    int y = r - (r - l) / 3;

    if (func(x) < func(y)) {
        l = x;
    } else {
        r = y;
    }

    for (int i=l; i<=r; i++) {
        ans = max(ans, func(i));
    }
}

// segtree iterative

const int N = 1e5;
int t[2*N];
int n;

void build () {
    for (int i=n-1; i>0; i--) {
        t[i] = t[i<<1] + t[i<<1|1];
    }
}

void modify (int i, int x) {
    i += n;
    t[i] = x;
    while (i>1) {
        t[i>>1] = t[i] + t[i^1];
        i>>=1;
    }
}

int query (int l, int r) {
    l += n; r += n;
    int ans = 0;
    while (l<r) {
        if (l&1) ans += t[l++]; 
        if (r&1) ans += t[--r];
        l >>= 1; r >>= 1;
    }
    return ans
}