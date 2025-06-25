import { createRouter, createWebHistory } from 'vue-router'
import Login from '../components/Login.vue'
import Dashboard from '../components/dashboard.vue'
import MyLearning from '../components/MyLearning.vue'
import Courses from '../components/Courses.vue'

const routes = [
	{ path: '/', component: Login },
	{
		path: '/dashboard',
		component: Dashboard,
		children: [
			{ path: 'my-learning', component: MyLearning },
			{ path: 'courses', component: Courses },
		],
	},
]

const router = createRouter({
	history: createWebHistory(),
	routes,
})

export default router
