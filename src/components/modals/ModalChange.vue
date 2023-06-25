<script>
import {mapActions} from "vuex";

export default {
  data() {
    return {
      name: "ModalChange",
      synthes: null,
      newNucleotides: null,
      visible: false,
      error: null
    }
  },
  methods: {
    ...mapActions({
      "delete": "synthesizer/deleteSynthesisMethod",
      "updateNucleotides": "synthesizer/updateNucleotidesMethod",
      "updateStatus": "synthesizer/updateStatusMethod",
      "updateTimer": "synthesizer/updateTimerMethod",
    }),
    changeNucleotides() {
      if (!this.synthes.work) {
        const data = {
          nucleotides: this.newNucleotides.toLowerCase().split(""),
          synthes: this.synthes
        }
        this.updateNucleotides(data)
        this.updateStatus("Бездействует")
        this.updateTimer()
        this.visible = false
      }
    },
    deleteSynthes() {
      this.delete(this.synthes)
      this.updateStatus("Бездействует")
      this.updateTimer()
      this.visible = false
    },
    openModal(synthes) {
      this.synthes = synthes
      this.newNucleotides = synthes.nucleotides.join("")
      this.updateStatus("Редактирование")
      this.visible = true
    }
  },
  computed: {
    isChanged() {
      const regular = /^([a, t, g, c]+|\d+)$/i;
      if (!this.newNucleotides) {
        this.error = null
        return false
      }

      if (!regular.test(this.newNucleotides)) {
        this.error = "Введите корректную формулу!"
        return true
      }

      if (this.newNucleotides.length < 6) {
        this.error = "Слишком короткая формула!"
        return true
      }

      if (this.newNucleotides.length > 120) {
        this.error = "Слишком длинная формула!"
        return true
      }

      if (!this.newNucleotides || !regular.test(this.newNucleotides)) {
        return true
      }
      return false
    }
  }
}
</script>

<template>
  <div v-show="visible" class="modal">
    <div class="modal--content">
      <p v-show="isChanged" class="error--formula">{{ this.error }}</p>
      <input v-model="newNucleotides" placeholder="Введите формулу" class="synthesis__description error--border" />
      <div class="synthesis__btns">
        <button :disabled="isChanged" @click="changeNucleotides()" class="synthesis__btn medium">Изменить</button>
        <button :disabled="isChanged" @click="deleteSynthes()" class="synthesis__btn high">Удалить</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.error--formula {
  padding: 10px 20px;
  border-radius: 10px;
  top: -3em;
  background-color: #d73838;
  color: #fff;
}
.modal {
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  z-index: 1;
  background: rgba(0,0,0,0.3);
  display: flex;
  align-items: center;
  transition: all 0.3s ease-in;
  border-radius: 6px;
}

.modal--content {
  width: 50%;
  display: flex;
  position: relative;
  margin: 0 auto;
  border-radius: 6px;
}

.synthesis__description {
  border-width: 0;
}
</style>