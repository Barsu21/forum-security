package main

import (
	"crypto/tls"
	"fmt"
	"forum/database"
	"forum/handlers"
	"forum/handlers/auth"
	"forum/handlers/comment"
	"forum/handlers/others"
	"forum/handlers/post"
	"net/http"
)

func main() {
	const (
		CertFilePath = "./tls/cert.pem"
		KeyFilePath  = "./tls/key.pem"
	)

	database.InitDB()

	rateLimiter := handlers.NewRateLimiter()

	h := routes()

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      ":8080",
		Handler:   rateLimiter.LimitMiddleware(h),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running at https://localhost:8080")
	err := server.ListenAndServeTLS(CertFilePath, KeyFilePath)
	if err != nil {
		fmt.Println("Error starting HTTPS server:", err)
	}
}
func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))
	mux.Handle("/web/images/", http.StripPrefix("/web/images/", http.FileServer(http.Dir("./web/images"))))
	mux.HandleFunc("/", post.MainPage)
	mux.HandleFunc("/sign_in", auth.SignIn)
	mux.HandleFunc("/sign_up", auth.SignUp)
	mux.HandleFunc("/sign_out", auth.SignOut)
	mux.HandleFunc("/create_post", post.CreatePost)
	mux.HandleFunc("/my_posts", post.MyPosts)
	mux.HandleFunc("/posts/", post.Post)
	mux.HandleFunc("/category/", post.Filter)
	mux.HandleFunc("/react", post.React)
	mux.HandleFunc("/liked_posts", post.LikedPosts)
	mux.HandleFunc("/delete_post", post.DeletePost)
	mux.HandleFunc("/comment", comment.Comment)
	mux.HandleFunc("/react_comment", comment.ReactComment)
	mux.HandleFunc("/delete_comment", comment.DeleteComment)
	mux.HandleFunc("/login", auth.GoogleLogin)
	mux.HandleFunc("/auth/callback", auth.GoogleCallback)
	mux.HandleFunc("/login_github", auth.GithubLogin)
	mux.HandleFunc("/auth/github/callback", auth.GitHubCallback)
	mux.HandleFunc("/user_profile", others.Profile)
	mux.HandleFunc("/apply", others.Apply)
	mux.HandleFunc("/admin_page", others.AdminPage)
	mux.HandleFunc("/admin_page_allow", others.ModeratorAllow)
	mux.HandleFunc("/admin_page_deny", others.ModeratorDeny)
	mux.HandleFunc("/admin_page_demote_to_user", others.DemoteToUser)
	mux.HandleFunc("/report_to_admin", others.ReportToAdmin)
	mux.HandleFunc("/send_reply", others.SendReply)
	mux.HandleFunc("/delete_report_from_admin", others.DeleteReportFromAdminPage)
	mux.HandleFunc("/delete_reply_from_admin", others.DeleteReplyByID)
	return mux
}
