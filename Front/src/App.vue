<script setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
	<router-view />
</template>

<script>
import Login from './components/Login.vue'

export default {
	components: {
		Login,
	},
}
</script>
<style scoped>
.logo {
	height: 6em;
	padding: 1.5em;
	will-change: filter;
	transition: filter 300ms;
}
.logo:hover {
	filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
	filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
