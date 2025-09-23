const Dashboard = {
    template: `
        <div>
            <div class="header">
                <div class="logo">仪表盘</div>
                <div class="user-info">
                    <span>{{ this.$root.username }}</span>
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
                    <label class="control-label">路由保护开关</label>
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
            protectionStatus: true,
            routeProtect: true,
            ethernetProtect: true,
            wirelessProtect: true,
            onlineDevices: 5
        }
    },
    methods: {
        logout() {
            this.$router.push('/login');
        },
        refreshData() {
            // 从/api/status获取数据
            fetch('/api/status')
                .then(response => response.json())
                .then(data => {
                    this.onlineDevices = data.client;
                    this.routeProtect = data['route-protect'];
                    this.ethernetProtect = data['ethernet-protect'];
                    this.wirelessProtect = data['wireless-protect'];
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
                    token: localStorage.getItem('token'),
                    'route-protect': this.routeProtect
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.status === 0) {
                    alert('操作失败');
                } else {
                    refreshData();
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
                    token: localStorage.getItem('token'),
                    'ethernet-protect': this.ethernetProtect
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.status === 0) {
                    alert('操作失败');
                } else {
                    refreshData();
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
                    token: localStorage.getItem('token'),
                    'wireless-protect': this.wirelessProtect
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.status === 0) {
                    alert('操作失败');
                } else {
                    refreshData();
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