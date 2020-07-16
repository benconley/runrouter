<template>
  <div style="height: 500px; width: 100%">
    <l-map
      ref="map"
      :zoom="zoom"
      :center="center"
      :options="mapOptions"
      style="height: 80%"
      @click="addStartMarker"
      @update:center="centerUpdate"
      @update:zoom="zoomUpdate"
    >
      <l-tile-layer
        :url="url"
        :attribution="attribution"
      />
      <l-marker :lat-lng="startMarker.latlng" v-if="startMarker"></l-marker>
      <l-polyline ref="features" :lat-lngs="nodeCoords"></l-polyline>
    </l-map>
  </div>
</template>

<script>
import { latLng } from "leaflet";
import { LMap, LTileLayer, LMarker, LPolyline } from "vue2-leaflet";

export default {
  name: "Example",
  components: {
    LMap,
    LTileLayer,
    LMarker,
    LPolyline
  },
  props: [
      "latlngCallback",
      "nodeCoords",
  ],
  data() {
    return {
      zoom: 13,
      center: latLng(36.012493, -78.966372),
      url: 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',
      attribution:
        '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
      currentZoom: 11.5,
      currentCenter: latLng(36.012493, -78.966372),
      mapOptions: {
        zoomSnap: 0.5
      },
      startMarker: null,
    };
  },
  methods: {
    addStartMarker(e) {
      this.startMarker = {latlng: e.latlng}
      this.latlngCallback(e.latlng)
    },
    zoomUpdate(zoom) {
      this.currentZoom = zoom;
    },
    centerUpdate(center) {
      this.currentCenter = center;
    },
  },
  watch: {
    nodeCoords() {
      this.$nextTick(() => {
        this.$refs.map.mapObject.fitBounds(this.$refs.features.mapObject.getBounds())
      })
    }
  }
};
</script>
