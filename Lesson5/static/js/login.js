const API = 'http://localhost:8080'

async function Login() {
    const body = {
        login: document.getElementById("input_login").value,
        password: document.getElementById("input_password").value
    }
    try {
        const res = await fetch(`${API}/login`, {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(body)
        })
        if (!res.ok){
            showMessage("login_msg", "Неверный логин или пароль.", "error")
            return
        }
        const data = await res.json()
        localStorage.setItem("token", data.token)
        showMessage("login_msg", "Успешный вход. Перенаправляем...", "success")
        setTimeout(() =>{
            window.location.href = "/users.html"
        }, 1000)
    } catch (e) {
        showMessage("login_msg", "Сервер недоступен.", "error")
    }
}

function showMessage(id, text, type) {
    const el = document.getElementById(id)
    el.textContent = text
    el.className = `message ${type}`
}