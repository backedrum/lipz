<style scoped>
  p {
    margin-left: 25px;
  }
  #closeHolder {
    text-align: center;
  }
  #close {
    background-color: mistyrose;
  }
</style>
<template>
  <div>
    <capture-settings/>
    <modal name="warn"
      :width="300"
      :height="100">
      <p>Traffic on this device cannot be captured. Please choose another one.</p>
      <div id="closeHolder">
      <button id="close" @click="$modal.hide('warn')">Close</button>
      </div>
    </modal>
    <vue-good-table
      title="Devices Available"
      :columns="columns"
      :rows="rows"
      theme="nocturnal"
      @on-row-click="onRowClickFn">
    <div slot="emptystate">
      Sorry, there is nothing to display.
    </div>
    </vue-good-table>
  </div>
</template>

<script>
  import CaptureSettings from './CaptureSettings.vue'

  export default {
    components: {CaptureSettings},
    name: 'devicesTable',
    data () {
      return {
        columns: [
          {
            label: 'Name',
            field: 'name'
          },
          {
            label: 'Address IPv4',
            field: 'addressIP4'
          },
          {
            label: 'Address IPv6',
            field: 'addressIP6'
          }
        ],
        rows: []
      }
    },
    methods: {
      getTableData () {
        var self = this

        setTimeout(function () {
          self.$http.get('/api/devices/list')
            .then(response => {
              for (var i = 0; i < response.body.Devices.length; i++) {
                var device = response.body.Devices[i]
                var addressIP6 = device.Addresses.length > 0
                  ? device.Addresses[0].IP : '-'
                var addressIP4 = device.Addresses.length > 1
                  ? device.Addresses[1].IP : '-'

                self.rows.push({'name': device.Name, 'addressIP4': addressIP4, 'addressIP6': addressIP6})
              }
            }, response => {
              console.log(response)
            })
        }, 100)
      },
      onRowClickFn (params) {
        if (params.row.addressIP4 === '-') {
          this.$modal.show('warn', {
            title: 'Cannot capture!'
          })
          return
        }

        this.$modal.show('captureSettings', {
          title: 'Start capture',
          interfaceName: params.row.name,
          ip4Address: params.row.addressIP4,
          ip6Address: params.row.addressIP6})
      }
    },
    created () {
      this.getTableData()
    }
  }
</script>
