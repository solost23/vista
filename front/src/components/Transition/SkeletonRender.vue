<template>
  <template v-if="!isInit">
    <slot key="fake" name="fake"></slot>
  </template>
  <template v-else>
    <slot key="content" name="content"></slot>
  </template>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'SkeletonRender',
  props: {
    isInit: {
      type: Boolean,
      default: false
    }
  }
})
</script>
