<script lang="ts" setup>
const count = await fetch("/api/count")
    .then((res) => res.json())
    .then((data) => data.count);
</script>

<template>
    <p>This page has been visited {{ count }} times!</p>
</template>
