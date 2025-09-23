const EthernetRules = {
    template: `
        <div>
            <div class="header">
                <div class="logo">有线网络管理</div>
                <div class="user-info">
                    <span>管理员</span>
                    <button class="btn" @click="logout">退出</button>
                </div>
            </div>
            
            <div class="card">
                <div class="card-title">有线网络保护设置</div>
                <div class="form-group">
                    <label class="control-label">有线网络保护开关</label>
                    <label class="toggle-switch">
                        <input type="checkbox" v-model="ethernetProtect">
                        <span class="slider"></span>
                    </label>
                    <span style="margin-left: 10px;">{{ ethernetProtect ? '已启用' : '已禁用' }}</span>
                </div>
                <button class="btn" @click="saveSettings">保存设置</button>
            </div>
        </div>
    `,
    data() {
        return {
            ethernetProtect: true
        }
    },
    methods: {
        saveSettings() {
            // 保存设置到后端 /api/modify/ethernet-protect
            fetch('/api/modify/ethernet-protect', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    token: localStorage.getItem('token'),
                    'ethernet-protect': this.ethernetProtect
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.status === 1) {
                    fetch('/api/rules')
                    .then(response => response.json())
                    .then(data => {
                        this.ethernetProtect = data['ethernet-protect'];
                    })
                    .catch(error => {
                        console.error('Error fetching rules:', error);
                    });
                    alert('设置已保存');
                } else {
                    alert('保存失败');
                }
            })
            .catch(error => {
                console.error('Error saving settings:', error);
                alert('保存失败');
            });
        },
        logout() {
            this.$router.push('/login');
        }
    },
    mounted() {
        // 获取初始状态
        fetch('/api/rules')
            .then(response => response.json())
            .then(data => {
                this.ethernetProtect = data['ethernet-protect'];
            })
            .catch(error => {
                console.error('Error fetching rules:', error);
            });
    }
};