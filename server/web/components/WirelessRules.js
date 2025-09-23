const WirelessRules = {
    template: `
        <div>
            <div class="header">
                <div class="logo">无线网络管理</div>
                <div class="user-info">
                    <span>{{ username }}</span>
                    <button class="btn" @click="logout">退出</button>
                </div>
            </div>
            
            <div class="card">
                <div class="card-title">无线网络保护设置</div>
                <div class="form-group">
                    <label class="control-label">无线网络保护开关</label>
                    <label class="toggle-switch">
                        <input type="checkbox" v-model="wirelessProtect" @change="saveSettings">
                        <span class="slider"></span>
                    </label>
                    <span style="margin-left: 10px;">{{ wirelessProtect ? '已启用' : '已禁用' }}</span>
                </div>
                <button class="btn" @click="saveSettings">保存设置</button>
            </div>
            
            <div class="card">
                <div class="card-title">添加无线网络规则</div>
                <div class="form-group">
                    <label>SSID</label>
                    <input type="text" v-model="newRule.ssid" placeholder="例如：OfficeWiFi">
                </div>
                <div class="form-group">
                    <label>描述</label>
                    <input type="text" v-model="newRule.description" placeholder="例如：办公区WiFi">
                </div>
                <button class="btn" @click="addRule">添加规则</button>
            </div>
            
            <div class="card">
                <div class="card-title">无线网络规则列表</div>
                <table class="rules-table">
                    <thead>
                        <tr>
                            <th>SSID</th>
                            <th>描述</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(rule, index) in rules" :key="index">
                            <td>{{ rule.ssid }}</td>
                            <td>{{ rule.description }}</td>
                            <td>
                                <button class="btn btn-danger" @click="deleteRule(index)">删除</button>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    `,
    data() {
        return {
            username: localStorage.getItem('username'),
            wirelessProtect: true,
            newRule: {
                ssid: '',
                description: ''
            },
            rules: []
        }
    },
    methods: {
        saveSettings() {
            // 保存设置到后端 /api/modify/wireless-protect
            fetch('/api/modify/wireless-protect', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Token: localStorage.getItem('token'),
                    'Wireless_Protect': this.wirelessProtect
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.status === 1) {
                    alert('设置已保存');
                } else {
                    alert('操作失败');
                }
            })
            .catch(error => {
                console.error('Error saving settings:', error);
                alert('操作失败');
            });
        },
        addRule() {
            if (this.newRule.ssid) {
                const wirelessRules = [...this.rules, { ...this.newRule }];
                this.saveWirelessRules(wirelessRules);
                this.newRule = { ssid: '', description: '' };
            } else {
                alert('请填写SSID');
            }
        },
        deleteRule(index) {
            const wirelessRules = [...this.rules];
            wirelessRules.splice(index, 1);
            this.saveWirelessRules(wirelessRules);
        },
        saveWirelessRules(wirelessRules) {
            fetch('/api/modify/wireless-rules', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Token: localStorage.getItem('token'),
                    'Wireless_Rules': wirelessRules.map(rule => ({
                        'SSID': rule.ssid,
                        'Description': rule.description
                    }))
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.status === 1) {
                    this.rules = wirelessRules;
                    alert('保存成功');
                } else {
                    alert('操作失败');
                }
            })
            .catch(error => {
                console.error('Error saving wireless rules:', error);
                alert('操作失败');
            });
        },
        logout() {
            fetch('/api/auth/logout', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Token: localStorage.getItem('token')
                })
            }).catch(error => {
                console.error('Error logging out:', error);
            });
            localStorage.setItem('username', '');
            localStorage.setItem('token', '');
            this.$router.push('/login');
        },
        fetchWirelessRules() {
            fetch('/api/rules')
                .then(response => response.json())
                .then(data => {
                    if (data['Wireless_Rules']) {
                        this.rules = data['Wireless_Rules'].map(rule => ({
                            ssid: rule.SSID,
                            description: rule.Description
                        }));
                    }
                    this.wirelessProtect = data['Wireless_Protect'] || true;
                })
                .catch(error => {
                    console.error('Error fetching wireless rules:', error);
                });
        }
    },
    mounted() {
        this.fetchWirelessRules();
    }
};