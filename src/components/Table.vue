<script>
import IconArrowLeft from "@/components/icons/IconArrowLeft.vue";
import IconArrowRight from "@/components/icons/IconArrowRight.vue";
import {mapActions} from "vuex";
import ModalChange from "@/components/modals/ModalChange.vue";

export default {
  name: "Table",
  data() {
    return {
      newNucleotides: null,
      visible: false,
      visibleForm: false
    }
  },
  components: {ModalChange, IconArrowLeft, IconArrowRight},
  props: ['synthesis', 'title', 'leftPriority', 'rightPriority'],
  methods: {
    ...mapActions({
      "update": "synthesizer/updatePriorityMethod"
    }),
    updatePriority(synthes, futurePriority) {
      if (!synthes.work) {
        let data = {
          futurePriority: futurePriority,
          synthes: synthes
        }
        this.update(data)
      }
    },
    showChangeModal(synthes, ref) {
      if(!synthes.work) {
        this.$refs.change.openModal(synthes)
      }
    },
  },
}
</script>

<template>
  <div class="task">
    <span class="task--title">{{ title }}</span>
    <ModalChange ref="change"/>
    <div v-for="synthes in synthesis" class="task--body">
      <div :class="synthes.work ? 'synthes--work synthes--block' : 'synthes--block'">
        <button @click="updatePriority(synthes, leftPriority)" class="btn--priority">
          <IconArrowLeft/>
        </button>

        <span @click="showChangeModal(synthes)">{{ synthes.nucleotides.join("") }}</span>

        <button @click="updatePriority(synthes, rightPriority)" class="btn--priority">
          <IconArrowRight/>
        </button>
      </div>
    </div>
  </div>
</template>
