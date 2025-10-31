<template>
    <div class="auth-container">
        <div class="auth-card">
            <h2>{{ isLogin ? '用户登录' : '用户注册' }}</h2>
            
            <form @submit.prevent="handleSubmit" class="auth-form">
                <!-- 账号ID输入框 -->
                <div class="form-group" v-if="isLogin">
                    <label>账号ID：</label>
                    <input 
                        type="number" 
                        v-model="formData.account_id"
                        placeholder="请输入您的账号ID"
                        required
                    />
                </div>

                <!-- 用户名输入框 -->
                <div class="form-group" v-if="!isLogin">
                    <label>用户名：</label>
                    <input 
                        type="text" 
                        v-model="formData.username"
                        placeholder="请输入用户名（3-20位）"
                        required
                    />
                </div>

                <!-- 密码输入框 -->
                <div class="form-group">
                    <label>密码：</label>
                    <input 
                        type="password" 
                        v-model="formData.password"
                        :placeholder="isLogin ? '请输入密码' : '请输入密码（至少6位）'"
                        required
                    />
                </div>

                <!-- 邮箱输入框 -->
                <div class="form-group" v-if="!isLogin">
                    <label>邮箱：</label>
                    <input 
                        type="email" 
                        v-model="formData.email"
                        placeholder="请输入邮箱地址"
                        required
                    />
                </div>

                <!-- 昵称输入框 -->
                <div class="form-group" v-if="!isLogin">
                    <label>昵称：</label>
                    <input 
                        type="text" 
                        v-model="formData.nickname"
                        placeholder="请输入昵称（可选）"
                    />
                </div>

                <button type="submit" class="submit-btn" :disabled="loading">
                    {{ loading ? '处理中...' : (isLogin ? '登录' : '注册') }}
                </button>
            </form>

            <div class="auth-switch">
                <span>{{ isLogin ? '没有账号？' : '已有账号？' }}</span>
                <button type="button" @click="$emit('toggle-mode')" class="switch-btn">
                    {{ isLogin ? '立即注册' : '立即登录' }}
                </button>
            </div>

            <!-- 消息提示 -->
            <div v-if="message" :class="['message', messageType]">
                {{ message }}
            </div>
        </div>
    </div>
</template>

<script>
import { ref } from 'vue'
import { register, login } from '../api/auth.js'

export default {
    props: {
        isLogin: {
            type: Boolean,
            default: true
        }
    },
    emits: ['toggle-mode', 'login-success'],
    setup(props, { emit }) {
        const formData = ref({
            account_id: '',
            username: '',
            password: '',
            email: '',
            nickname: ''
        })
        const loading = ref(false)
        const message = ref('')
        const messageType = ref('info')

        const showMessage = (text, type = 'info') => {
            message.value = text
            messageType.value = type
        }

        const resetForm = () => {
            formData.value = {
                account_id: '',
                username: '',
                password: '',
                email: '',
                nickname: ''
            }
        }

        const handleSubmit = async () => {
            loading.value = true
            message.value = ''

            try {
                if (props.isLogin) {
                    // 登录逻辑
                    const result = await login(
                        formData.value.account_id,
                        formData.value.password
                    )
                    showMessage(`登录成功！欢迎 ${result.data.username}`, 'success')
                    emit('login-success', result.data)
                    
                } else {
                    // 注册逻辑
                    const result = await register(formData.value)
                    showMessage(
                        `注册成功！您的账号ID是：${result.data.account_id}，请牢记！`,
                        'success'
                    )
                    
                    // 注册成功后自动切换到登录模式
                    emit('toggle-mode')
                }
                
                resetForm()
                
            } catch (error) {
                showMessage(error.message, 'error')
            } finally {
                loading.value = false
            }
        }

        return {
            formData,
            loading,
            message,
            messageType,
            handleSubmit,
            showMessage
        }
    }
}
</script>

<style scoped>
.auth-container {
    width: 100%;
    max-width: 400px;
    padding: 20px;
}

.auth-card {
    background: white;
    padding: 40px 30px;
    border-radius: 10px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}

h2 {
    text-align: center;
    margin-bottom: 30px;
    color: #333;
}

.form-group {
    margin-bottom: 20px;
}

label {
    display: block;
    margin-bottom: 5px;
    color: #555;
    font-weight: 500;
}

input {
    width: 100%;
    padding: 12px;
    border: 2px solid #e1e5e9;
    border-radius: 6px;
    font-size: 14px;
    transition: border-color 0.3s;
}

input:focus {
    outline: none;
    border-color: #667eea;
}

.submit-btn {
    width: 100%;
    padding: 12px;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 16px;
    cursor: pointer;
    transition: opacity 0.3s;
}

.submit-btn:hover:not(:disabled) {
    opacity: 0.9;
}

.submit-btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.auth-switch {
    text-align: center;
    margin-top: 20px;
    color: #666;
}

.switch-btn {
    background: none;
    border: none;
    color: #667eea;
    cursor: pointer;
    text-decoration: underline;
}

.message {
    margin-top: 15px;
    padding: 10px;
    border-radius: 4px;
    text-align: center;
    font-size: 14px;
}

.message.success {
    background: #d4edda;
    color: #155724;
    border: 1px solid #c3e6cb;
}

.message.error {
    background: #f8d7da;
    color: #721c24;
    border: 1px solid #f5c6cb;
}

.message.info {
    background: #d1ecf1;
    color: #0c5460;
    border: 1px solid #bee5eb;
}
</style>