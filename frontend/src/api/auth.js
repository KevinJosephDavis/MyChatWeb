const API_BASE = 'http://localhost:8080/api';

// 统一的请求函数
async function request(url, options = {}) {
    try {
        const response = await fetch(`${API_BASE}${url}`, {
            headers: {
                'Content-Type': 'application/json',
                ...options.headers,
            },
            ...options,
            body: options.body ? JSON.stringify(options.body) : undefined,
        });
        
        const data = await response.json();
        
        if (!response.ok) {
            throw new Error(data.message || '请求失败');
        }
        
        return data;
    } catch (error) {
        throw new Error(`网络请求失败: ${error.message}`);
    }
}

// 用户注册
export async function register(userData) {
    return request('/register', {
        method: 'POST',
        body: userData,
    });
}

// 用户登录
export async function login(accountID, password) {
    return request('/login', {
        method: 'POST',
        body: {
            account_id: parseInt(accountID),
            password: password,
        },
    });
}