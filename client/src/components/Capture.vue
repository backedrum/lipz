<template>
  <vue-good-table
    title="Captured packets"
    :columns="columns"
    :rows="rows"
    theme="nocturnal"
    :pagination-options="{
      enabled: true,
      rowsPerPageLabel: 'Packets per page'
    }"
    :globalSearch="true"
    :search-options="{
      enabled: true
    }"
    :sort-options="{
      enabled: true
    }"
  >
    <div slot="emptystate">
      <vue-simple-spinner line-fg-color="#666666"></vue-simple-spinner>
    </div>
  </vue-good-table>
</template>

<script>
  import VueSimpleSpinner from 'vue-simple-spinner'

  export default {
    components: {VueSimpleSpinner},
    name: 'packetsTable',
    data () {
      return {
        columns: [
          {
            label: 'Protocol',
            field: 'protocol'
          },
          {
            label: 'Src IP',
            field: 'srcIP'
          },
          {
            label: 'Src port',
            field: 'srcPort'
          },
          {
            label: 'Dst IP',
            field: 'dstIP'
          },
          {
            label: 'Dst port',
            field: 'dstPort'
          },
          {
            label: 'Dump',
            field: 'dump',
            width: '500px'
          },
          {
            label: 'Payload',
            field: 'payload',
            width: '150px'
          }
        ],
        rows: []
      }
    },
    methods: {
      getTableData () {
        var self = this

        setTimeout(function () {
          self.$http.post('/api/capture/' +
            self.$route.params.interfaceName,
            {
              'duration': parseInt(self.$route.params.duration, 10),
              'filename': self.$route.params.filename
            },
            {
              headers: {
                'Content-Type': 'application/json'
              }
            })
            .then(response => {
              for (var i = 0; i < response.body.NetPacketInfoList.length; i++) {
                var packet = response.body.NetPacketInfoList[i]
                self.rows.push({
                  'protocol': packet.Protocol,
                  'srcIP': packet.Src,
                  'srcPort': packet.SrcPort,
                  'dstIP': packet.Dst,
                  'dstPort': packet.DstPort,
                  'dump': packet.Dump,
                  'payload': packet.Payload
                })
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
