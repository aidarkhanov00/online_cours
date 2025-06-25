<template>
	<aside class="sidebar">
		<div class="logo">
			<img src="../assets/icons/logo.png" alt="Логотип" />
		</div>
		<ul>
			<li>
				<router-link to="/dashboard/my-learning">
					<img src="../assets/icons/my-learning.png" alt="Мое обучение" /> Мое
					обучение
				</router-link>
			</li>
			<li>
				<router-link to="/dashboard/courses">
					<img src="../assets/icons/disciplines.png" alt="Все предметы" /> Все
					предметы
				</router-link>
			</li>
			<li>
				<router-link to="/dashboard/materials">
					<img src="../assets/icons/materials.png" alt="Материалы" /> Материалы
				</router-link>
			</li>
			<li>
				<router-link to="/dashboard/notifications">
					<img src="../assets/icons/notifications.png" alt="Уведомления" />
					Уведомления
				</router-link>
			</li>
			<li>
				<router-link to="/dashboard/settings">
					<img src="../assets/icons/settings.png" alt="Настройки" /> Настройки
				</router-link>
			</li>
		</ul>
		<button @click="logout" class="logout-btn">
			<img src="../assets/icons/logout.png" alt="Выход" /> Выйти
		</button>
	</aside>
</template>

<script>
import { useRouter } from 'vue-router'

export default {
	setup() {
		const router = useRouter()
		const logout = () => {
			localStorage.removeItem('token') // Удаляем токен
			router.push('/') // Перенаправляем на страницу входа
		}
		return { logout }
	},
}
</script>

<style scoped>
.sidebar {
	width: 170px;
	background: #2c3e50;
	color: white;
	padding: 20px;
	height: 100vh;
	display: flex;
	flex-direction: column;
	justify-content: center; /* Центрируем содержимое */
	align-items: center;
}

/* Логотип сверху */
.logo {
	width: 100%;
	display: flex;
	justify-content: center;
	margin-bottom: 20px;
}

.sidebar .logo img {
	width: 50px; /* Размер иконки */
	height: 50px;
}

.sidebar ul {
	list-style: none;
	width: 100%;
}

.sidebar li {
	margin: 15px 0;
	display: flex;
	align-items: center;
}

.sidebar a {
	color: white;
	text-decoration: none;
	font-size: 16px;
	display: flex;
	align-items: center;
	gap: 10px; /* Отступ между иконкой и текстом */
	width: 100%;
	padding: 10px;
	border-radius: 5px;
	transition: background 0.3s;
}

.sidebar a:hover {
	background: #34495e;
}

.sidebar img {
	width: 24px; /* Размер иконки */
	height: 24px;
}

.logout-btn {
	margin-top: 20px;
	padding: 10px;
	background: #e74c3c;
	color: white;
	border: none;
	border-radius: 5px;
	cursor: pointer;
	width: 70%;
	display: flex;
	align-items: center;
	justify-content: center;
	gap: 10px;
	font-size: 16px;
}

.logout-btn img {
	width: 20px;
	height: 20px;
}

.logout-btn:hover {
	background: #c0392b;
}
</style>
