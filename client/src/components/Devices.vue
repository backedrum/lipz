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
      :paginate="true"
      :lineNumbers="true"
      :onClick="onClickFn">
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
            field: 'name',
            filterable: true
          },
          {
            label: 'Address IPv4',
            field: 'addressIP4',
            filterable: true
          },
          {
            label: 'Address IPv6',
            field: 'addressIP6',
            filterable: true
          }
        ],
        rows: [],
        onClickFn: function (row, index) {
          if (row.addressIP4 === '-') {
            this.$modal.show('warn', {
              title: 'Cannot capture!'
            })
            return
          }

          this.$modal.show('captureSettings', {title: 'Start capture', interfaceName: row.name, ip4Address: row.addressIP4, ip6Address: row.addressIP6})
        }
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
      }
    },
    created () {
      this.getTableData()
    }
  }
</script>
