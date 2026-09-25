package main

import (
	"fmt"
	"log"
	"net/http"
)

const htmlPage = `<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <title>Gabriel | Desenvolvedor</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            font-family: 'Segoe UI', sans-serif;
            background: linear-gradient(-45deg, #0a192f, #0f3d3e, #123456, #064635);
            background-size: 400% 400%;
            animation: gradientShift 12s ease infinite;
            color: #e0f2f1;
        }

        @keyframes gradientShift {
            0% { background-position: 0% 50%; }
            50% { background-position: 100% 50%; }
            100% { background-position: 0% 50%; }
        }

        .card {
            background: rgba(255, 255, 255, 0.05);
            backdrop-filter: blur(10px);
            border: 1px solid rgba(255, 255, 255, 0.1);
            border-radius: 16px;
            padding: 48px 56px;
            text-align: center;
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
            max-width: 420px;
        }

        h1 {
            font-size: 2rem;
            margin-bottom: 8px;
            background: linear-gradient(90deg, #4fd1c5, #63b3ed);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        p {
            font-size: 1rem;
            color: #b0c4c0;
            margin-bottom: 24px;
        }

        .tag {
            display: inline-block;
            padding: 6px 16px;
            margin: 4px;
            border-radius: 999px;
            background: rgba(79, 209, 197, 0.15);
            border: 1px solid rgba(79, 209, 197, 0.4);
            font-size: 0.85rem;
            color: #4fd1c5;
        }
    </style>
</head>
<body>
    <div class="card">
        <h1>Gabriel</h1>
        <p>Desenvolvedor Backend de Go e explorando infraestrutura cloud da Oracle.</p>
        <span class="tag">Go</span>
        <span class="tag">Backend</span>
        <span class="tag">Cloud</span>
    </div>
</body>-
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlPage)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Servidor rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
