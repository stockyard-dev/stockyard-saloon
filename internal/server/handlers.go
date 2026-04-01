package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-saloon/internal/store")
func(s *Server)handleListBoards(w http.ResponseWriter,r *http.Request){list,_:=s.db.ListBoards();if list==nil{list=[]store.Board{}};writeJSON(w,200,list)}
func(s *Server)handleCreateBoard(w http.ResponseWriter,r *http.Request){var b store.Board;json.NewDecoder(r.Body).Decode(&b);if b.Name==""{writeError(w,400,"name required");return};if err:=s.db.CreateBoard(&b);err!=nil{writeError(w,500,err.Error());return};writeJSON(w,201,b)}
func(s *Server)handleDeleteBoard(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.DeleteBoard(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleListThreads(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListThreads(id);if list==nil{list=[]store.Thread{}};writeJSON(w,200,list)}
func(s *Server)handleCreateThread(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var t store.Thread;json.NewDecoder(r.Body).Decode(&t);t.BoardID=id;if t.Title==""{writeError(w,400,"title required");return};if err:=s.db.CreateThread(&t);err!=nil{writeError(w,500,err.Error());return};writeJSON(w,201,t)}
func(s *Server)handleDeleteThread(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.DeleteThread(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleListPosts(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListPosts(id);if list==nil{list=[]store.Post{}};writeJSON(w,200,list)}
func(s *Server)handleCreatePost(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var p store.Post;json.NewDecoder(r.Body).Decode(&p);p.ThreadID=id;if p.Body==""{writeError(w,400,"body required");return};if err:=s.db.CreatePost(&p);err!=nil{writeError(w,500,err.Error());return};writeJSON(w,201,p)}
func(s *Server)handleDeletePost(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.DeletePost(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleStats(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
