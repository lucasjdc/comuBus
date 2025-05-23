
const loginTab = document.getElementById('loginTab');
const signupTab = document.getElementById('signupTab');
const loginForm = document.getElementById('loginForm');
const signupForm = document.getElementById('signupForm');

loginTab.addEventListener('click', () => {
    loginTab.classList.add('active');
    signupTab.classList.remove('active');
    loginForm.classList.add('active');
    signupForm.classList.remove('active');
});

signupTab.addEventListener('click', () => {
    signupTab.classList.add('active');
    loginTab.classList.remove('active');
    signupForm.classList.add('active');
    loginForm.classList.remove('active');
});

document.getElementById('forgotPasswordLink').addEventListener('click', (e) => {
    e.preventDefault();
    alert('Instruções para recuperação de senha serão enviadas ao seu e-mail.');
});

loginForm.addEventListener('submit', (e) => {
    e.preventDefault();
    alert('Login efetuado (simulado).');
    loginForm.reset();
});

signupForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const data = {
        nome: document.getElementById('signupName').value,
        email: document.getElementById('signupEmail').value,
        senha: document.getElementById('signupPassword').value, // Corrigido aqui
        cidade: document.getElementById('signupCity').value,
        linhas_favoritas: Array.from(document.querySelectorAll('input[name="favLines"]:checked')).map(cb => cb.value)
    };

    try {
        const response = await fetch('/cadastro', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        const result = await response.json();

        if (response.ok) {
            alert(result.mensagem || 'Cadastro realizado com sucesso!');
            signupForm.reset();

            // Volta para aba de login
            loginTab.classList.add('active');
            signupTab.classList.remove('active');
            loginForm.classList.add('active');
            signupForm.classList.remove('active');
        } else {
            alert(result.mensagem || 'Erro no cadastro');
        }
    } catch (error) {
        alert('Erro na comunicação com o servidor.');
        console.error(error);
    }
});