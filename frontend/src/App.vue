<template>
    <div id="app">
        <div class="app-header" v-if="user">
            <h1>聊天室测试界面</h1>
            <div class="user-info">
                欢迎，{{ user.username }}！
                <button @click="logout" class="logout-btn">退出登录</button>
            </div>
        </div>
        
        <div class="app-content">
            <AuthForm 
                v-if="!user"
                :is-login="isLoginMode" 
                @toggle-mode="toggleAuthMode"
                @login-success="handleLoginSuccess"
            />
            
            <div v-else class="welcome-message">
                <h2>登录成功！</h2>
                <p>用户ID: {{ user.user_id }}</p>
                <p>账号ID: {{ user.account_id }}</p>
                <p>用户名: {{ user.username }}</p>
                <p>接下来可以开始开发聊天功能了！</p>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import AuthForm from './components/AuthForm.vue'
import { useAuthStore } from './stores/auth.js'

export default {
    name: 'App',
    components: {
        AuthForm
    },
    setup() {
        const isLoginMode = ref(true)
        const { user, setUser, logout: authLogout } = useAuthStore()

        const toggleAuthMode = () => {
            isLoginMode.value = !isLoginMode.value
        }

        const handleLoginSuccess = (userData) => {
            setUser(userData)
        }

        const logout = () => {
            authLogout()
        }

        onMounted(() => {
            // 检查本地存储中的登录状态
            const savedUser = localStorage.getItem('user')
            if (savedUser) {
                setUser(JSON.parse(savedUser))
            }
        })

        return {
            isLoginMode,
            user,
            toggleAuthMode,
            handleLoginSuccess,
            logout
        }
    }
}
</script>

<style>
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: 'Microsoft YaHei', sans-serif;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
}

#app {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
}

.app-header {
    background: rgba(255, 255, 255, 0.1);
    padding: 20px;
    color: white;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.app-content {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 20px;
}

.welcome-message {
    background: white;
    padding: 40px;
    border-radius: 10px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
    text-align: center;
    max-width: 500px;
}

.welcome-message h2 {
    color: #4CAF50;
    margin-bottom: 20px;
}

.welcome-message p {
    margin: 10px 0;
    color: #666;
}

.logout-btn {
    background: #ff4757;
    color: white;
    border: none;
    padding: 8px 16px;
    border-radius: 4px;
    cursor: pointer;
    margin-left: 15px;
}

.logout-btn:hover {
    background: #ff3742;
}

.user-info {
    display: flex;
    align-items: center;
}
</style>