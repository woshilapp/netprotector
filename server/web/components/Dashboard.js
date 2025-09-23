const Dashboard = {
    template: `
        <div>
            <div class="header">
                <div class="logo">仪表盘</div>
                <div class="user-info">
                    <span>{{ username }}</span>
                    <button class="btn" @click="logout">退出</button>
                </div>
            </div>
            
            <div class="dashboard">
                <div class="card">
                    <div class="card-title">保护状态</div>
                    <div class="stat-number">
                        <span class="status-indicator" :class="protectionStatus ? 'status-active' : 'status-inactive'"></span>
                        {{ protectionStatus ? '已启用' : '已禁用' }}
                    </div>
                    <div class="stat-label">当前防护状态</div>
                </div>
                
                <div class="card">
                    <div class="card-title">在线设备</div>
                    <div class="stat-number">{{ onlineDevices }}</div>
                    <div class="stat-label">当前连接设备数</div>
                </div>
            </div>
            
            <div class="controls">
                <div class="control-group">
                    <label class="control-label">总保护开关(请勿关闭)</label>
                    <label class="toggle-switch">
                        <input type="checkbox" v-model="routeProtect" @change="saveRouteProtect">
                        <span class="slider"></span>
                    </label>
                    <span style="margin-left: 10px;">{{ routeProtect ? '启用中' : '已禁用' }}</span>
                </div>
                
                <div class="control-group">
                    <label class="control-label">有线网络保护</label>
                    <label class="toggle-switch">
                        <input type="checkbox" v-model="ethernetProtect" @change="saveEthernetProtect">
                        <span class="slider"></span>
                    </label>
                    <span style="margin-left: 10px;">{{ ethernetProtect ? '启用中' : '已禁用' }}</span>
                </div>
                
                <div class="control-group">
                    <label class="control-label">无线网络保护</label>
                    <label class="toggle-switch">
                        <input type="checkbox" v-model="wirelessProtect" @change="saveWirelessProtect">
                        <span class="slider"></span>
                    </label>
                    <span style="margin-left: 10px;">{{ wirelessProtect ? '启用中' : '已禁用' }}</span>
                </div>
                
                <div class="control-group">
                    <label class="control-label">操作</label>
                    <button class="btn" @click="refreshData">刷新数据</button>
                </div>
            </div>
        </div>
    `,
    data() {
        return {
            username: localStorage.getItem('username'),
            protectionStatus: true,
            routeProtect: true,
            ethernetProtect: true,
            wirelessProtect: true,
            onlineDevices: 5
        }
    },
    methods: {
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
        refreshData() {
            // 从/api/status获取数据
            fetch('/api/status', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Token: localStorage.getItem('token')
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.Status !== 1) {
                    this.logout();
                    return;
                }
                this.onlineDevices = data.Client;
                this.routeProtect = data['Route_Protect'];
                this.ethernetProtect = data['Ethernet_Protect'];
                this.wirelessProtect = data['Wireless_Protect'];
            })
            .catch(error => {
                console.error('Error fetching status:', error);
            });
        },
        saveRouteProtect() {
            fetch('/api/modify/route-protect', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Token: localStorage.getItem('token'),
                    'Route_Protect': this.routeProtect
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.Status === 0) {
                    alert('操作失败');
                } else {
                    this.refreshData();
                }
            })
            .catch(error => {
                console.error('Error saving route protect:', error);
                alert('操作失败');
            });
        },
        saveEthernetProtect() {
            fetch('/api/modify/ethernet-protect', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Token: localStorage.getItem('token'),
                    'Ethernet_Protect': this.ethernetProtect
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.Status === 0) {
                    alert('操作失败');
                } else {
                    this.refreshData();
                }
            })
            .catch(error => {
                console.error('Error saving ethernet protect:', error);
                alert('操作失败');
            });
        },
        saveWirelessProtect() {
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
                if (data.Status === 0) {
                    alert('操作失败');
                } else {
                    this.refreshData();
                }
            })
            .catch(error => {
                console.error('Error saving wireless protect:', error);
                alert('操作失败');
            });
        }
    },
    mounted() {
        this.refreshData();
    }
};