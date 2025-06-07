document.addEventListener('DOMContentLoaded', () => {
    const registerForm = document.getElementById('registerForm');
    const loginForm = document.getElementById('loginForm');
    const addServerForm = document.getElementById('addServerForm');
    const listServersBtn = document.getElementById('listServersBtn');

    const registerMessage = document.getElementById('registerMessage');
    const loginMessage = document.getElementById('loginMessage');
    const addServerMessage = document.getElementById('addServerMessage');
    const serverListDiv = document.getElementById('serverList');

    // Registration Logic
    if (registerForm) {
        registerForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const username = document.getElementById('regUsername').value;
            const password = document.getElementById('regPassword').value;

            try {
                const response = await fetch('/register', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ username, password })
                });
                const data = await response.json();
                if (response.ok) {
                    registerMessage.textContent = data.message;
                    registerMessage.className = 'success';
                } else {
                    registerMessage.textContent = data.message || 'Registration failed';
                    registerMessage.className = 'error';
                }
            } catch (error) {
                registerMessage.textContent = 'Error: ' + error.message;
                registerMessage.className = 'error';
            }
        });
    }

    // Login Logic
    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const username = document.getElementById('loginUsername').value;
            const password = document.getElementById('loginPassword').value;

            try {
                const response = await fetch('/login', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ username, password })
                });
                const data = await response.json();
                if (response.ok) {
                    loginMessage.textContent = data.message;
                    loginMessage.className = 'success';
                    // In a real app, store token/session here
                } else {
                    loginMessage.textContent = data.message || 'Login failed';
                    loginMessage.className = 'error';
                }
            } catch (error) {
                loginMessage.textContent = 'Error: ' + error.message;
                loginMessage.className = 'error';
            }
        });
    }

    // Add Server Logic
    if (addServerForm) {
        addServerForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const name = document.getElementById('serverName').value;
            const ip = document.getElementById('serverIp').value;

            try {
                const response = await fetch('/add_server', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ name, ip })
                });
                const data = await response.json();
                if (response.ok) {
                    addServerMessage.textContent = data.message;
                    addServerMessage.className = 'success';
                    document.getElementById('serverName').value = '';
                    document.getElementById('serverIp').value = '';
                } else {
                    addServerMessage.textContent = data.message || 'Failed to add server';
                    addServerMessage.className = 'error';
                }
            } catch (error) {
                addServerMessage.textContent = 'Error: ' + error.message;
                addServerMessage.className = 'error';
            }
        });
    }

    // List Servers Logic
    if (listServersBtn) {
        listServersBtn.addEventListener('click', async () => {
            try {
                const response = await fetch('/list_servers');
                const servers = await response.json();

                if (response.ok) {
                    if (servers.length === 0) {
                        serverListDiv.innerHTML = '<p>No servers added yet.</p>';
                        return;
                    }
                    const ul = document.createElement('ul');
                    servers.forEach(server => {
                        const li = document.createElement('li');
                        li.textContent = `Name: ${server.name}, IP: ${server.ip}`;
                        ul.appendChild(li);
                    });
                    serverListDiv.innerHTML = ''; // Clear previous list
                    serverListDiv.appendChild(ul);
                } else {
                    serverListDiv.innerHTML = `<p class="error">Error: ${servers.message || 'Failed to fetch servers'}</p>`;
                }
            } catch (error) {
                serverListDiv.innerHTML = `<p class="error">Error: ${error.message}</p>`;
            }
        });
    }
});
