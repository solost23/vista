<script setup>
// setup: 组合式写法，自动完成注册组件
import Header from "./components/Header.vue"
import Footer from "./components/Footer.vue"
</script>

<template>
    <Header></Header>
    <RouterView></RouterView>
    <Footer></Footer>
</template>
