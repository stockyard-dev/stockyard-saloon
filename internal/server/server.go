package server
import("encoding/json";"net/http";"github.com/stockyard-dev/stockyard-saloon/internal/store")
type Server struct{db *store.DB;limits Limits;mux *http.ServeMux}
func New(db *store.DB,tier string)*Server{s:=&Server{db:db,limits:LimitsFor(tier),mux:http.NewServeMux()};s.routes();return s}
func(s *Server)ListenAndServe(addr string)error{return(&http.Server{Addr:addr,Handler:s.mux}).ListenAndServe()}
func(s *Server)routes(){
    s.mux.HandleFunc("GET /health",s.handleHealth)
    s.mux.HandleFunc("GET /api/stats",s.handleStats)
    s.mux.HandleFunc("GET /api/boards",s.handleListBoards)
    s.mux.HandleFunc("POST /api/boards",s.handleCreateBoard)
    s.mux.HandleFunc("DELETE /api/boards/{id}",s.handleDeleteBoard)
    s.mux.HandleFunc("GET /api/boards/{id}/threads",s.handleListThreads)
    s.mux.HandleFunc("POST /api/boards/{id}/threads",s.handleCreateThread)
    s.mux.HandleFunc("DELETE /api/threads/{id}",s.handleDeleteThread)
    s.mux.HandleFunc("GET /api/threads/{id}/posts",s.handleListPosts)
    s.mux.HandleFunc("POST /api/threads/{id}/posts",s.handleCreatePost)
    s.mux.HandleFunc("DELETE /api/posts/{id}",s.handleDeletePost)
    s.mux.HandleFunc("GET /",s.handleUI)
}
func(s *Server)handleHealth(w http.ResponseWriter,r *http.Request){writeJSON(w,200,map[string]string{"status":"ok","service":"stockyard-saloon"})}
func writeJSON(w http.ResponseWriter,status int,v interface{}){w.Header().Set("Content-Type","application/json");w.WriteHeader(status);json.NewEncoder(w).Encode(v)}
func writeError(w http.ResponseWriter,status int,msg string){writeJSON(w,status,map[string]string{"error":msg})}
func(s *Server)handleUI(w http.ResponseWriter,r *http.Request){if r.URL.Path!="/"{http.NotFound(w,r);return};w.Header().Set("Content-Type","text/html");w.Write(dashboardHTML)}
