const RouteRules = {
    template: `
        <div>
            <div class="header">
                <div class="logo">路由规则管理</div>
                <div class="user-info">
                    <span>{{ username }}</span>
                    <button class="btn" @click="logout">退出</button>
                </div>
            </div>
            
            <div class="card">
                <div class="card-title">添加路由规则</div>
                <div class="form-group">
                    <label>网络地址</label>
                    <input type="text" v-model="newRule.network" placeholder="例如：192.168.1.0">
                </div>
                <div class="form-group">
                    <label>子网掩码</label>
                    <input type="text" v-model="newRule.mask" placeholder="例如：255.255.255.0">
                </div>
                <div class="form-group">
                    <label>网关</label>
                    <input type="text" v-model="newRule.endpoint" placeholder="例如：192.168.1.1">
                </div>
                <div class="form-group">
                    <label>描述</label>
                    <input type="text" v-model="newRule.description" placeholder="例如：办公区网络">
                </div>
                <button class="btn" @click="addRule">添加规则</button>
            </div>
            
            <div class="card">
                <div class="card-title">路由规则列表</div>
                <table class="rules-table">
                    <thead>
                        <tr>
                            <th>网络地址</th>
                            <th>子网掩码</th>
                            <th>网关</th>
                            <th>描述</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(rule, index) in rules" :key="index">
                            <td>{{ rule.network }}</td>
                            <td>{{ rule.mask }}</td>
                            <td>{{ rule.endpoint }}</td>
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
            newRule: {
                network: '',
                mask: '',
                endpoint: '',
                description: ''
            },
            rules: []
        }
    },
    methods: {
        addRule() {
            if (this.newRule.network && this.newRule.mask && this.newRule.endpoint) {
                const routeRules = [...this.rules, { ...this.newRule }];
                this.saveRouteRules(routeRules);
                this.newRule = { network: '', mask: '', endpoint: '', description: '' };
            } else {
                alert('请填写必要信息');
            }
        },
        deleteRule(index) {
            const routeRules = [...this.rules];
            routeRules.splice(index, 1);
            this.saveRouteRules(routeRules);
        },
        saveRouteRules(routeRules) {
            fetch('/api/modify/route-rules', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Token: localStorage.getItem('token'),
                    'Route_Rules': routeRules.map(rule => ({
                        'Network': rule.network,
                        'Mask': rule.mask,
                        'Endpoint': rule.endpoint,
                        'Description': rule.description
                    }))
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.Status === 1) {
                    this.rules = routeRules;
                    this.fetchRouteRules();
                    // alert('保存成功');
                } else {
                    alert('操作失败');
                }
            })
            .catch(error => {
                console.error('Error saving route rules:', error);
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
        fetchRouteRules() {
            fetch('/api/rules')
                .then(response => response.json())
                .then(data => {
                    if (data['Route_Rules']) {
                        this.rules = data['Route_Rules'].map(rule => ({
                            network: rule.Network,
                            mask: rule.Mask,
                            endpoint: rule.Endpoint,
                            description: rule.Description
                        }));
                    }
                })
                .catch(error => {
                    console.error('Error fetching route rules:', error);
                });
        }
    },
    mounted() {
        this.fetchRouteRules();
    }
};