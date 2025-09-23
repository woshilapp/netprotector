const Login = {
    template: `
        <div class="login-page">
            <div class="card" style="max-width: 400px; margin: 100px auto;">
                <div class="card-title">登录</div>
                <div class="form-group">
                    <label>用户名</label>
                    <input type="text" v-model="username" placeholder="请输入用户名">
                </div>
                <div class="form-group">
                    <label>密码</label>
                    <input type="password" v-model="password" placeholder="请输入密码">
                </div>
                <button class="btn" @click="login">登录</button>
            </div>
        </div>
    `,
    data() {
        return {
            username: '',
            password: ''
        }
    },
    methods: {
        login() {
            fetch('/api/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Username: this.username,
                    Password: this.password
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.Status === 1) {
                    localStorage.setItem('token', data.Token);
                    localStorage.setItem('username', this.username);
                    this.$router.push('/dashboard');
                } else {
                    alert('登录失败');
                }
            })
            .catch(error => {
                console.error('Error logging in:', error);
                alert('登录失败');
            });
        }
    }
};