<template>
	<div class="login-container">
		<h2>Вход</h2>
		<input v-model="username" type="text" placeholder="Имя пользователя" />
		<input v-model="password" type="password" placeholder="Пароль" />
		<button @click="login">Войти</button>
		<p v-if="errorMessage" class="error">{{ errorMessage }}</p>
	</div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api'

export default {
	setup() {
		const username = ref('')
		const password = ref('')
		const router = useRouter() // Получаем объект роутера

		const login = async () => {
			try {
				const response = await api.post('/login', {
					username: username.value,
					password: password.value,
				})
				localStorage.setItem('token', response.data.token) // Сохраняем токен
				router.push('/dashboard') // Переход на dashboard
			} catch (error) {
				console.error('Ошибка авторизации:', error)
			}
		}

		return { username, password, login }
	},
}
</script>

<style scoped>
.login-container {
	max-width: 300px;
	margin: auto;
	padding: 20px;
	border: 1px solid #ccc;
	border-radius: 8px;
	text-align: center;
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
	background: white; /* На случай, если фон не белый */
	box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}
input {
	width: 70%; /* Сделал шире, чтобы выглядело лучше */
	padding: 10px;
	margin: 10px 0;
	border: 1px solid #ccc;
	border-radius: 5px;
}
button {
	width: 70%;
	padding: 10px;
	background: #42b983;
	color: white;
	border: none;
	border-radius: 5px;
	cursor: pointer;
}
button:hover {
	background: #369a77;
}
</style>
