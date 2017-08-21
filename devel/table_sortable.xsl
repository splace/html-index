<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:template match="base">
<html> <script src="sorttable.js"></script>
<body>
  <h2>Index Of http(s)://<xsl:value-of select="@host"/><xsl:value-of select="@name"/></h2>
 <h3><a href="..">Parent Directory</a></h3>
  <table sortable="true" class="sortable">
 	<thead>
 	<tr>
 	<th style="width:1em"></th>
  	<th>Name</th>
 	<th style="width:10%;">Size</th>
	<th style="width:30%">Last Modified</th>
 	</tr>
 	</thead>
 	<tbody> 
  
    <xsl:for-each select="*[local-name()='file' or local-name()='dir' or local-name()='raw']">
    <tr>
    <td>
    <xsl:choose>
    <xsl:when test="local-name()='dir'">
    <xsl:attribute name="style">border-style:double solid solid solid;border-width:4px 1px 1px 1px;border-radius: 10%;</xsl:attribute>
    </xsl:when>

     <xsl:when test="local-name()='file'">
   <xsl:attribute name="style">border:1px solid;border-top-right-radius:25%;</xsl:attribute>
	<table>
 	<thead>
 	<tr>
	<th style="width:1em;"></th>
 	</tr>
 	</thead> 	<tbody> 
	<tr><td style="border-style:double none double none"></td></tr>
 	</tbody> 
	</table> 
  </xsl:when>
   </xsl:choose>
  </td>
   
      <td><a><xsl:attribute name="href"><xsl:value-of select="@name"/></xsl:attribute><xsl:value-of select="@name"/></a></td>
      <td style="text-align:right"><xsl:value-of select="@size"/></td>
      <td><xsl:value-of select="@modified"/></td>
    </tr>
    </xsl:for-each>
 	</tbody> 
  </table>
</body>
</html>
</xsl:template>
</xsl:stylesheet>



