SYNO_VERSION=$(cat $SYNOLOGY_VERSION_FILE_PATH | grep -Po '(?<=productversion=").*(?=")')
SYNO_PRETTY_VERSION=$(echo $SYNO_VERSION | grep -Po '\d+.\d+')

echo NAME=\"Synology DSM\" > $OS_RELEASE_FILE_PATH
echo ID=dsm >> $OS_RELEASE_FILE_PATH
echo VERSION=$SYNO_VERSION >> $OS_RELEASE_FILE_PATH
echo PRETTY_NAME=\"Synology DSM v$SYNO_PRETTY_VERSION\" >> $OS_RELEASE_FILE_PATH
echo HOME_URL=\"https://synology.com\" >> $OS_RELEASE_FILE_PATH
echo BUG_REPORT_URL=\"https://synology.com\" >> $OS_RELEASE_FILE_PATH

cat $OS_RELEASE_FILE_PATH