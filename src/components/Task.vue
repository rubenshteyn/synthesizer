<script>
import ModalChange from "@/components/modals/ModalChange.vue";

export default {
  name: "Task",
  data() {
    return {
      visible: false
    }
  },
  props: ['synthes', 'queue'],
  components: {
    ModalChange
  },
  methods: {
    showChangeModal(synthes) {
      if (!synthes.work) {
        this.$refs.change.openModal(synthes, this.queue)
      }
    },
  }
}
</script>

<template>
  <ModalChange ref="change"/>
  <div v-if="synthes.nucleotides.join('').length > 17" class="task--block">

    <div v-if="!visible" class="task__item">
      <div @click="showChangeModal(synthes)">{{ synthes.nucleotides.join("").substr(0, 17) }}</div>
    </div>
    <div v-if="!visible"  @click="visible = true" class="disclosure">...</div>

    <div v-if="visible" class="task__item">
      <div @click="showChangeModal(synthes)" v-for="nucleotide in synthes.nucleotides">
        {{ nucleotide }}
      </div>
    </div>
    <div v-if="visible" @click="visible = false" class="collapse">
      свернуть
    </div>
  </div>
  <div v-else class="task__item" @click="showChangeModal(synthes)">
    <span>{{ synthes.nucleotides.join("") }}</span>
  </div>
</template>

<style scoped>

</style>