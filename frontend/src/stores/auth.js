import { ref } from 'vue'

export const useAuthStore = () => {
    const user = ref(null)
    const isAuthenticated = ref(false)

    const setUser = (userData) => {
        user.value = userData
        isAuthenticated.value = true
        localStorage.setItem('user', JSON.stringify(userData))
    }

    const logout = () => {
        user.value = null
        isAuthenticated.value = false
        localStorage.removeItem('user')
    }

    // 检查本地存储中的登录状态
    const checkAuth = () => {
        const savedUser = localStorage.getItem('user')
        if (savedUser) {
            user.value = JSON.parse(savedUser)
            isAuthenticated.value = true
        }
    }

    return {
        user,
        isAuthenticated,
        setUser,
        logout,
        checkAuth
    }
}