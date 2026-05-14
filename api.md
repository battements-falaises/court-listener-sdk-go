# Courts

Response Types:

- <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Court">Court</a>

Methods:

- <code title="get /courts/{id}/">client.Courts.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#CourtService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#CourtGetParams">CourtGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Court">Court</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /courts/">client.Courts.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#CourtService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#CourtListParams">CourtListParams</a>) (\*<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go/packages/pagination#CursorURLPage">CursorURLPage</a>[<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Court">Court</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Dockets

Response Types:

- <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Docket">Docket</a>

Methods:

- <code title="get /dockets/{id}/">client.Dockets.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#DocketService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#DocketGetParams">DocketGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Docket">Docket</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /dockets/">client.Dockets.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#DocketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#DocketListParams">DocketListParams</a>) (\*<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go/packages/pagination#CursorURLPage">CursorURLPage</a>[<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Docket">Docket</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Clusters

Response Types:

- <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Cluster">Cluster</a>

Methods:

- <code title="get /clusters/{id}/">client.Clusters.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#ClusterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#ClusterGetParams">ClusterGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Cluster">Cluster</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clusters/">client.Clusters.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#ClusterService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#ClusterListParams">ClusterListParams</a>) (\*<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go/packages/pagination#CursorURLPage">CursorURLPage</a>[<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Cluster">Cluster</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Opinions

Response Types:

- <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Opinion">Opinion</a>

Methods:

- <code title="get /opinions/{id}/">client.Opinions.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#OpinionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#OpinionGetParams">OpinionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Opinion">Opinion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /opinions/">client.Opinions.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#OpinionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#OpinionListParams">OpinionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go/packages/pagination#CursorURLPage">CursorURLPage</a>[<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go">courtlistenersdk</a>.<a href="https://pkg.go.dev/github.com/battements-falaises/court-listener-sdk-go#Opinion">Opinion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
