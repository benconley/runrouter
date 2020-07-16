<template>
  <div id="app">
    <Form :submitCallback="sendLatLng" :latitude="latitude" :longitude="longitude" />
    <Map :latlngCallback="setLatLng" :nodeCoords="nodeCoords"/>
  </div>
</template>

<script>
import Form from './components/Form.vue'
import Map from './components/Map.vue'

export default {
  name: 'App',
  components: {
    Form,
    Map
  },
  data() {
    return {
      latitude: '',
      longitude: '',
      nodeCoords: [],
    };
  },
  methods: {
    sendLatLng(payload) {
      console.log(payload)
      fetch('http://localhost:8080/route', {
          method: 'POST',
          body: JSON.stringify(payload),
        })
        .then(response => response.json())
        .then(data => {
          console.log('Success:', data);
          this.setNodeCoords(data)
        })
        .catch((error) => {
          console.error('Error:', error);
      });
    },
    setLatLng(latlng) {
      this.latitude = latlng.lat
      this.longitude = latlng.lng
    },
    setNodeCoords(payload) {
      this.nodeCoords = payload.nodes.map(node => {
        return [node.latitude, node.longitude]
      })
    }
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
