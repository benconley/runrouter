<template>
  <div>
    <div id="form-container">
      <span v-if="latitudeError">{{ latitudeError }}</span>
      <span v-if="longitudeError">{{ longitudeError }}</span>
      <label for="distance-input">
        Distance
        <input id="distance-input" v-model="distanceInput" type="number">
        <span v-if="distanceError">{{ distanceError }}</span>
      </label>
    </div>
    <div id="buttons-container">
      <button id="submit-button" @click="submitInput">
        Submit
      </button>
    </div>
  </div>
</template>

<script>

export default {
  props: [
      "submitCallback",
      "latitude",
      "longitude"
  ],
  data() {
    return {
      distanceInput: '',
      latitudeError: '',
      longitudeError: '',
      distanceError: '',
    };
  },
  methods: {
    submitInput() {
      this.latitudeError = null
      this.longitudeError = null
      this.distanceError = null

      if (this.latitude === '') {
        this.latitudeError = 'Latitude cannot be null.'
      }
      if (this.longitude === '') {
        this.longitudeError = 'Longitude cannot be null.'
      }
      if (this.distanceInput === '' || this.distanceInput <= 0) {
        this.distanceError = 'Distance must be greater than zero.'
      }

      if (this.longitudeError == null && this.latitudeError == null && this.distanceError == null) {
        console.log('done')
        const payload = {
          start: {
            latitude: this.latitude,
            longitude: this.longitude
          },
          distance: parseFloat(this.distanceInput)
        }
        this.submitCallback(payload)
      }
    }
  },
}
</script>

<style scoped>
#input {
  text-align: center;
  margin: 60px 0 0 0;
}
</style>
