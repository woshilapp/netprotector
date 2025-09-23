const TimeRules = {
    template: `
        <div>
            <div class="header">
                <div class="logo">时间规则管理</div>
                <div class="user-info">
                    <span>管理员</span>
                    <button class="btn" @click="logout">退出</button>
                </div>
            </div>
            
            <div class="card">
                <div class="card-title">添加时间规则</div>
                <div class="form-group">
                    <label>规则名称</label>
                    <input type="text" v-model="newRule.name" placeholder="例如：晚上禁网">
                </div>
                <div class="form-group">
                    <label>开始时间</label>
                    <input type="time" v-model="newRule.startTime">
                </div>
                <div class="form-group">
                    <label>结束时间</label>
                    <input type="time" v-model="newRule.endTime">
                </div>
                <button class="btn" @click="addRule">添加规则</button>
            </div>
            
            <div class="card">
                <div class="card-title">时间规则列表</div>
                <table class="rules-table">
                    <thead>
                        <tr>
                            <th>规则名称</th>
                            <th>开始时间</th>
                            <th>结束时间</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(rule, index) in rules" :key="index">
                            <td>{{ rule.name }}</td>
                            <td>{{ rule.startTime }}</td>
                            <td>{{ rule.endTime }}</td>
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
            newRule: {
                name: '',
                startTime: '',
                endTime: ''
            },
            rules: []
        }
    },
    methods: {
        addRule() {
            if (this.newRule.name && this.newRule.startTime && this.newRule.endTime) {
                // 转换时间格式
                const timeRules = [...this.rules, { ...this.newRule }];
                this.saveTimeRules(timeRules);
                this.newRule = { name: '', startTime: '', endTime: '' };
            } else {
                alert('请填写完整信息');
            }
        },
        deleteRule(index) {
            const timeRules = [...this.rules];
            timeRules.splice(index, 1);
            this.saveTimeRules(timeRules);
        },
        saveTimeRules(timeRules) {
            fetch('/api/modify/time-rules', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    token: localStorage.getItem('token'),
                    'time-rules': timeRules.map(rule => ({
                        'time-start': rule.startTime,
                        'time-end': rule.endTime,
                        'description': rule.name
                    }))
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.status === 1) {
                    this.rules = timeRules;
                    fetchTimeRules();
                    alert('保存成功');
                } else {
                    alert('操作失败');
                }
            })
            .catch(error => {
                console.error('Error saving time rules:', error);
                alert('操作失败');
            });
        },
        logout() {
            this.$router.push('/login');
        },
        fetchTimeRules() {
            fetch('/api/rules')
                .then(response => response.json())
                .then(data => {
                    if (data['time-rules']) {
                        this.rules = data['time-rules'].map(rule => ({
                            name: rule.description,
                            startTime: rule['time-start'],
                            endTime: rule['time-end']
                        }));
                    }
                })
                .catch(error => {
                    console.error('Error fetching time rules:', error);
                });
        }
    },
    mounted() {
        this.fetchTimeRules();
    }
};