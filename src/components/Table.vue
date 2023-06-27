<script>
import IconArrowLeft from "@/components/icons/IconArrowLeft.vue";
import IconArrowRight from "@/components/icons/IconArrowRight.vue";
import {mapActions} from "vuex";
import ModalChange from "@/components/modals/ModalChange.vue";
import Task from "@/components/Task.vue";

export default {
  name: "Table",
  components: {Task, ModalChange, IconArrowLeft, IconArrowRight},
  props: ['synthesis', 'title', 'leftPriority', 'rightPriority', 'queue'],
  methods: {
    ...mapActions({
      "update": "synthesizer/updatePriorityMethod"
    }),
    updatePriority(synthes, futureQueue) {
      if (!synthes.work) {
        let data = {
          futureQueue: futureQueue,
          synthes: synthes,
          queue: this.queue
        }
        this.update(data)
      }
    },
  },
}
</script>

<template>
  <div class="task">
    <span class="task--title">{{ title }}</span>
    <div v-for="synthes in synthesis" class="task--body">
      <div :class="synthes.work ? 'synthes--work synthes--block' : 'synthes--block'">
        <button @click="updatePriority(synthes, leftPriority)" class="btn--priority">
          <IconArrowLeft/>
        </button>

        <Task :synthes="synthes" :queue="queue"/>

        <button @click="updatePriority(synthes, rightPriority)" class="btn--priority">
          <IconArrowRight/>
        </button>
      </div>
    </div>
  </div>
</template>

