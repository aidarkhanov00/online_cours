<template>
	<div class="main-dashboard">
		<img
			src="../assets/icons/illustration.png"
			alt="Иллюстрация"
			class="dashboard-image"
		/>
		<div class="text">
			<h1>Привет, Жандос</h1>
			<p>Рад(а) снова тебя видеть.</p>
		</div>
	</div>
</template>

<style scoped>
.main-dashboard {
	display: flex; /* Размещаем элементы в строку */
	align-items: center; /* Выравниваем по вертикали */
	padding: 20px;
	background-color: #f4f4f4;
	border-radius: 40px;
	box-shadow: 6px 6px 15px rgba(0, 0, 0, 0.2);
	position: relative; /* Для смещения картинки */
}

.dashboard-image {
	width: 150px; /* Настроить размер изображения */
	height: auto;
	position: absolute;
	left: 300px; /* Смещаем картинку левее */
}
</style>
