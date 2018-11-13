<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0"  xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:template match="index">
    <html>
		<head>
			<link rel="stylesheet" type="text/css" href="index.css"/>
		</head>
        <body>
       <table sortable="true">
            <thead>
                <tr>
	                <th colspan="4">
                		<h1>Index</h1>
						<h2>file://<xsl:value-of select="@host"/><xsl:value-of select="@name"/>/</h2>
						
					</th>
                </tr>
                <tr>
                    <th></th>
                    <th><h3>Name</h3></th>
                    <th><h3>Size</h3></th>
                    <th><h3>Last Modified</h3></th>
                </tr>
            </thead>
            <tbody>
				<tr><td>🗁</td><td><a href="..">.. [Parent Folder]</a></td></tr>
               <xsl:apply-templates/>
            </tbody>
        </table>
    </body>
</html>
</xsl:template>

<xsl:template match="dir">
    <tr>
        <td>🗀</td>
        <xsl:call-template name="completeRow"/>
    </tr>
</xsl:template>

<xsl:template match="txt">
    <tr>
	    <td>🖹</td>
         <xsl:call-template name="completeRow"/>
    </tr>
</xsl:template>

<xsl:template match="raw">
    <tr>
        <td/>
        <xsl:call-template name="completeRow"/>
    </tr>
</xsl:template>

<xsl:template name="completeRow">
    <td>
        <a>
            <xsl:attribute name="href">
                <xsl:value-of select="@name"/>
            </xsl:attribute>
            <xsl:value-of select="@name"/>
        </a>
    </td>
    <td>
        <xsl:value-of select="@size"/>
    </td>
    <td>
   <time>
		<xsl:attribute name="datetime">
		<xsl:value-of select="@modified"/>
		</xsl:attribute>
	    <xsl:value-of select="concat(substring(@modified, 1, 10 ),' ', substring(@modified, 12, 8 ))"/>
	</time>
    </td>
</xsl:template>

</xsl:stylesheet>


