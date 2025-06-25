<template>
	<div class="dashboard-layout">
		<Sidebar />

		<div class="content">
			<MainDashboard />
			<CourseCard />
		</div>

		<div class="tasks-column">
			<Header />
			<Tasks />
		</div>
	</div>
</template>

<script>
import Sidebar from '../components/Sidebar.vue'
import Header from '../components/Header.vue'
import MainDashboard from '../components/MainDashboard.vue'
import CourseCard from '../components/CourseCard.vue'
import Tasks from '../components/Tasks.vue'

export default {
	components: { Sidebar, Header, MainDashboard, CourseCard, Tasks },
}
</script>

<style scoped>
.dashboard-layout {
	display: flex;
	height: 100vh;
}

/* 1 Колонка: Sidebar */
.sidebar {
	width: 230px;
	background: black;
	color: white;
	padding: 20px;
	margin-top: 20px;
	height: 80vh;
	border-radius: 40px;
}

/* 2 Колонка: Dashboard + CourseCard */
.content {
	flex: 2;
	display: flex;
	flex-direction: column;
	padding: 20px;
	gap: 15px;
}

/* 3 Колонка: Header + Tasks */
.tasks-column {
	flex: 1;
	display: flex;
	flex-direction: column;
	padding: 20px;
	gap: 15px;
}
</style>
