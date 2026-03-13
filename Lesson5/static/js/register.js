const API = 'http://localhost:8080'

async function Register() {
    const body = {
        name: document.getElementById("input_name").value.trim(),
        age: parseInt(document.getElementById("input_age").value.trim()),
        login: document.getElementById("input_login").value.trim(),
        password: document.getElementById("input_password").value.trim(),
    }
    try {
        const res = await fetch(`${API}/register`, {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(body)
        })
        if (!res.ok) {
            showMessage("reg_msg", "Такой логин уже существует.", "error")
            return
        }
        const data = await res.json()
        showMessage("reg_msg", "Успешно, перенаправляем...", "success")
        setTimeout(() => {
           window.location.href = "/index.html" 
        }, 1000);
    }
    catch (e) {
        showMessage("reg_msg", "Сервер недоступен.", "error")
    }
}

function showMessage(id, text, type) {
    const el = document.getElementById(id)
    el.textContent = text
    el.className = `message ${type}`
}