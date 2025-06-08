# TerminalAI 🖥️🤖 
*Your personal AI assistant, right in your terminal.*

**TerminalAI** is a sleek, terminal-based AI assistant that brings Google's Gemini AI straight to your command line. Whether you're debugging code, seeking quick answers, or just want an AI chat companion, TerminalAI is built for speed, privacy, and productivity — right in your terminal.

---

## 💡 Features
- ⌨️ **Interactive terminal UI** powered by Bubble Tea
- 🧠 **Gemini API integration** for AI-powered conversations
- ⚡ **Fast and efficient** Golang backend
- 🔒 **Local-first experience** — minimal external dependencies

---

## 📊 HackaTime Progress  
![](https://hackatime-badge.hackclub.com/U07BWS2LPE2/TerminalAI)  
![](https://github-readme-stats.hackclub.dev/api/wakatime?username=2445&api_domain=hackatime.hackclub.com&theme=shadow_green&custom_title=Hackatime+Stats&layout=compact&cache_seconds=0&langs_count=8)
</br><img src="https://profile-counter.glitch.me/MadRobin13-TerminalAI/count.svg" alt="Visitor Counter"/>

---

## 🛠️ Tech Stack
**Frontend (Terminal UI)**:  
- Go Bubble Tea framework for terminal-based UI  

**Backend**:  
- Golang core application logic  
- Gemini API integration for AI responses  

**Database**:  
- Local file handling for direct access to previous chats (stored and labelled chronologically in the `./chats` folder)

---

## 🏁 Goal  
Built with ❤️ by [MadRobin13](https://github.com/MadRobin13) to make AI accessible, distraction-free, and integrated directly into your development workflow.

---

## 🚀 How to Use / Contribute  
1. Clone this repo:  
    ```bash
   git clone https://github.com/MadRobin13/TerminalAI.git
   cd TerminalAI
Install dependencies (Go modules):

2. Install Dependencies
    ```bash
    go get .
    go mod tidy

3. Create a .env file and add you gemini API key:
    ```bash
    API_KEY = {API_KEY}

4. Run TerminalAI:
    ```bash
    go run main.go

5. Fork, build, and send a PR to contribute!
