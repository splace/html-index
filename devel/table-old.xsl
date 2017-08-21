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
                		<h1 id="title">Folder Listing</h1>
						<h2>File://<xsl:value-of select="@host"/><xsl:value-of select="@name"/>/</h2>
					</th>
                </tr>
                <tr>
                    <th style="width:1em"></th>
                    <th style="width:50%"><h3>Name</h3></th>
                    <th><h3>Size</h3></th>
                    <th><h3>Last Modified</h3></th>
                </tr>
            </thead>
            <tbody>
				<tr><td></td><td><a href="..">.. [Parent Folder]</a></td><td></td><td></td></tr>
               <xsl:apply-templates/>
            </tbody>
        </table>
    </body>
</html>
</xsl:template>

<xsl:template match="dir">
    <tr>
        <xsl:call-template name="folderGraphic"/>
        <xsl:call-template name="completeRow"/>
    </tr>
</xsl:template>

<xsl:template match="file">
    <tr>
         <xsl:call-template name="fileGraphic"/>
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
    <td style="text-align:right">
        <xsl:value-of select="@size"/>
    </td>
    <td>
        <xsl:value-of select="@modified"/>
    </td>
</xsl:template>


<xsl:template name="folderGraphic">
    <td style="border-style:double solid solid solid;border-width:4px 1px 1px 1px;border-radius: 10%;">
    </td>
</xsl:template>

<xsl:template name="fileGraphic">
    <td style="border:1px solid;border-top-right-radius:25%;">
        <table>
            <th style="padding:0px;margin:0px">
                <tr>
                    <th style="padding:0px;margin:0px;width:1em;border-style:double none double none"/>
                </tr>
            </th>
        </table>
    </td>
</xsl:template>

</xsl:stylesheet>


