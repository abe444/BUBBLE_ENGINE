package functions

func DisplayHead() string {
return `<meta http-equiv="content-type" content="text/html; charset=utf-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    	<link rel="stylesheet" type="text/css" href="/static/main.css">
    	<meta name="description" content="Powered by the BUBBLE ENGINE.">`
}

func DisplayTable() string {
    return `<table>
               <tbody>
                <tr>
                    <td>
                    </td>
                    <td width=1000 valign="TOP">
                    <hr>
                    <hr>
                    </td>
                </tr>
                <tr>
                    <td valign="TOP">
                        <nav>
                        </nav>
                    </td>
                    <td height=500 valign="TOP">
                    </td>
                </tr>
                </tbody>
            </table>`
}
