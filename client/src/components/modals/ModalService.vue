<script setup>
import IconCogwheel from "@/components/icons/IconCogwheel.vue";
</script>

<template>
  <div class="modal" >
    <div class="modal--content">
      <span class="modal--title-top">СИНТЕЗАТ<IconCogwheel />Р</span>
      <span class="modal--title-bottom">НА <IconCogwheel />БСЛУЖИВАНИИ</span>
    </div>
  </div>
</template>

<style scoped>

.modal--cogwheel {
  position: relative;
  top: 7px;
}
.modal {
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  z-index: 1;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  transition: all 0.3s ease-in;
}

.modal--content {
  display: flex;
  flex-direction: column;
  position: relative;
  margin: 0 auto;
  color: #fff;
  font-size: 46px;
  padding: 3em 3em;
  background: rgb(64, 199, 129);;
  border-radius: 6px;
}
</style>