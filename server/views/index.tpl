<!DOCTYPE html>

<html>
<head>
    <title>MyDTs</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta http-equiv="refresh" content="30">
    <link rel="stylesheet" href="/static/css/style.css" type="text/css">
</head>
    <body>
        <h1>MDTs:</h1>
        <table class="mdt-table">
            <tr>
                <th>Serial Number</th>
                <th>Unit</th>
                <th>Unit ID</th>
                <th>Vehicle</th>
                <th>Signed On</th>
                <th>Internal IP</th>
                <th>Remote IP</th>
                <th>Updated</th>
            </tr>
            {{range .Mdts}}
                {{if .SignedOn}}
                    <tr class="signed-on">
                        <td>{{.SerialNumber}}</td>
                        <td>{{.UnitName}}</td>
                        <td>{{.UnitId}}</td>
                        <td>{{.VehicleId}}</td>
                        <td>{{.SignedOn}}</td>
                        <td>{{.InternalIp}}</td>
                        <td>{{.RemoteIp}}</td>
                        <td>{{.Updated}}</td>
                        <td><a href="vnc://{{.RemoteIp}}:5900" class="connect-button">Connect</a></td>
                    </tr>
                {{else}}
                    <tr class="signed-off">
                        <td>{{.SerialNumber}}</td>
                        <td>{{.UnitName}}</td>
                        <td>{{.UnitId}}</td>
                        <td>{{.VehicleId}}</td>
                        <td>{{.SignedOn}}</td>
                        <td>{{.InternalIp}}</td>
                        <td>{{.RemoteIp}}</td>
                        <td>{{.Updated}}</td>
                        <td><a href="vnc://{{.RemoteIp}}:5900" class="connect-button">Connect</a></td>
                    </tr>
                {{end}}
            {{end}}
        </table>
        <script src="/static/js/reload.min.js"></script>
    </body>
</html>