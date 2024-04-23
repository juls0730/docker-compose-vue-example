<script lang="ts" setup>
import { defineAsyncComponent } from "vue";

const VisitCount = defineAsyncComponent(
    () => import("../components/VisitCount.vue")
);
</script>

<template>
    <Suspense>
        <VisitCount />

        <template #fallback> Loading... </template>
    </Suspense>
</template>
