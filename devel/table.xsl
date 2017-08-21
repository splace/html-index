<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" version="1.0">
<xsl:template match="base">
    <html>
		<head>
			<link rel="stylesheet" type="text/css" href="table.css"/>
		</head>
        <body>
       <table sortable="true">
            <thead>
                <tr>
	                <th colspan="4">
                		<h1>Index</h1>
						<!--
						 <ul class="breadcrumb">
						  <li>file://</li>
						  <li><a href="#"><xsl:value-of select="@host"/></a></li>
						  <li><xsl:value-of select="@name"/></li>
						</ul> 	-->
						<h2>File://<xsl:value-of select="@host"/><xsl:value-of select="@name"/>/</h2>
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
				<tr><td>📂</td><td><a href="..">.. [Parent Folder]</a></td></tr>
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

<xsl:template match="file">
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
        <xsl:value-of select="@modified"/>
    </td>
</xsl:template>

</xsl:stylesheet>


