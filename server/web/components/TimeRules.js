const TimeRules = {
    template: `
        <div>
            <div class="header">
                <div class="logo">时间规则管理</div>
                <div class="user-info">
                    <span>{{ username }}</span>
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
                <div class="form-group">
                    <label>适用日期</label>
                    <div class="days-checkbox" style="display: flex; flex-wrap: wrap; gap: 10px; white-space: nowrap;">
                        <label v-for="(day, index) in days" :key="index" style="display: flex; align-items: center;">
                            <input type="checkbox" :value="day.value" v-model="newRule.days" style="margin-right: 5px;">
                            {{ day.label }}
                        </label>
                    </div>
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
                            <th>适用日期</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(rule, index) in rules" :key="index">
                            <td>{{ rule.name }}</td>
                            <td>{{ rule.startTime }}</td>
                            <td>{{ rule.endTime }}</td>
                            <td>{{ formatDays(rule.days) }}</td>
                            <td style="min-width: 200px;">
                            <div class="rule-days" style="display: flex; flex-wrap: wrap; gap: 10px; margin-top: 5px; white-space: nowrap;">
                                <label v-for="(day, dayIndex) in days" :key="dayIndex" style="display: flex; align-items: center;">
                                    <input type="checkbox" :value="day.value" v-model="rule.days" style="margin-right: 5px;">
                                    {{ day.label }}
                                </label>
                            </div>
                            </td>
                            <td>
                                <button class="btn btn-danger" @click="deleteRule(index)">删除</button>
                            </td>
                        </tr>
                    </tbody>
                </table>
                <div style="margin-top: 20px;">
                    <button class="btn" @click="saveAllRules">保存所有规则</button>
                </div>
            </div>
        </div>
    `,
    data() {
        return {
            username: localStorage.getItem('username'),
            newRule: {
                name: '',
                startTime: '',
                endTime: '',
                days: []
            },
            rules: [],
            days: [
                { value: 1, label: '周一' },
                { value: 2, label: '周二' },
                { value: 3, label: '周三' },
                { value: 4, label: '周四' },
                { value: 5, label: '周五' },
                { value: 6, label: '周六' },
                { value: 7, label: '周日' }
            ]
        }
    },
    methods: {
        addRule() {
            if (this.newRule.name && this.newRule.startTime && this.newRule.endTime) {
                // 转换时间格式
                this.rules.push({ ...this.newRule });
                this.newRule = { name: '', startTime: '', endTime: '', days: [] };
            } else {
                alert('请填写完整信息');
            }
        },
        deleteRule(index) {
            this.rules.splice(index, 1);
        },
        saveAllRules() {
            fetch('/api/modify/time-rules', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Token: localStorage.getItem('token'),
                    'Time_Rules': this.rules.map(rule => ({
                        'Time_Start': rule.startTime,
                        'Time_End': rule.endTime,
                        'Description': rule.name,
                        'Days': rule.days
                    }))
                })
            })
            .then(response => response.json())
            .then(data => {
                if (data.Status === 1) {
                    // alert('保存成功');
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
        fetchTimeRules() {
            fetch('/api/rules')
                .then(response => response.json())
                .then(data => {
                    if (data['Time_Rules']) {
                        this.rules = data['Time_Rules'].map(rule => ({
                            name: rule.Description,
                            startTime: rule['Time_Start'],
                            endTime: rule['Time_End'],
                            days: rule['Days'] || []
                        }));
                    }
                })
                .catch(error => {
                    console.error('Error fetching time rules:', error);
                });
        },
        formatDays(days) {
            if (!days || days.length === 0) return '每天';
            const dayLabels = this.days.filter(day => days.includes(day.value))
                                      .map(day => day.label);
            return dayLabels.join(', ');
        }
    },
    mounted() {
        this.fetchTimeRules();
    }
};