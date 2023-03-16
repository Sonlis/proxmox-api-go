package proxmox

import "strconv"

type QemuVirtIODisk struct {
	AsyncIO   QemuDiskAsyncIO   `json:"asyncio,omitempty"`
	Backup    bool              `json:"backup,omitempty"`
	Bandwidth QemuDiskBandwidth `json:"bandwidth,omitempty"`
	Cache     QemuDiskCache     `json:"cache,omitempty"`
	Discard   bool              `json:"discard,omitempty"`
	Format    QemuDiskFormat    `json:"format,omitempty"`
	Id        *uint             `json:"id,omitempty"`
	IOThread  bool              `json:"iothread,omitempty"`
	ReadOnly  bool              `json:"readonly,omitempty"`
	Replicate bool              `json:"replicate,omitempty"`
	Serial    QemuDiskSerial    `json:"serial,omitempty"`
	Size      uint              `json:"size,omitempty"`
	Storage   string            `json:"storage,omitempty"`
}

func (disk *QemuVirtIODisk) convertDataStructure() *qemuDisk {
	return &qemuDisk{
		AsyncIO:   disk.AsyncIO,
		Backup:    disk.Backup,
		Bandwidth: disk.Bandwidth,
		Cache:     disk.Cache,
		Discard:   disk.Discard,
		Disk:      true,
		Format:    disk.Format,
		Id:        disk.Id,
		IOThread:  disk.IOThread,
		ReadOnly:  disk.ReadOnly,
		Replicate: disk.Replicate,
		Serial:    disk.Serial,
		Size:      disk.Size,
		Storage:   disk.Storage,
		Type:      virtIO,
	}
}

func (disk QemuVirtIODisk) Validate() error {
	return disk.convertDataStructure().validate()
}

type QemuVirtIODisks struct {
	Disk_0  *QemuVirtIOStorage `json:"0,omitempty"`
	Disk_1  *QemuVirtIOStorage `json:"1,omitempty"`
	Disk_2  *QemuVirtIOStorage `json:"2,omitempty"`
	Disk_3  *QemuVirtIOStorage `json:"3,omitempty"`
	Disk_4  *QemuVirtIOStorage `json:"4,omitempty"`
	Disk_5  *QemuVirtIOStorage `json:"5,omitempty"`
	Disk_6  *QemuVirtIOStorage `json:"6,omitempty"`
	Disk_7  *QemuVirtIOStorage `json:"7,omitempty"`
	Disk_8  *QemuVirtIOStorage `json:"8,omitempty"`
	Disk_9  *QemuVirtIOStorage `json:"9,omitempty"`
	Disk_10 *QemuVirtIOStorage `json:"10,omitempty"`
	Disk_11 *QemuVirtIOStorage `json:"11,omitempty"`
	Disk_12 *QemuVirtIOStorage `json:"12,omitempty"`
	Disk_13 *QemuVirtIOStorage `json:"13,omitempty"`
	Disk_14 *QemuVirtIOStorage `json:"14,omitempty"`
	Disk_15 *QemuVirtIOStorage `json:"15,omitempty"`
}

// TODO write test
func (disks QemuVirtIODisks) mapToApiValues(currentDisks *QemuVirtIODisks, vmID uint, params map[string]interface{}, changes *qemuUpdateChanges) {
	tmpCurrentDisks := QemuVirtIODisks{}
	if currentDisks != nil {
		tmpCurrentDisks = *currentDisks
	}
	diskMap := disks.mapToIntMap()
	currentDiskMap := tmpCurrentDisks.mapToIntMap()
	for i := range diskMap {
		diskMap[i].convertDataStructure().markDiskChanges(currentDiskMap[i].convertDataStructure(), vmID, "virtio"+strconv.Itoa(int(i)), params, changes)
	}
}

func (disks QemuVirtIODisks) mapToIntMap() map[uint8]*QemuVirtIOStorage {
	return map[uint8]*QemuVirtIOStorage{
		0:  disks.Disk_0,
		1:  disks.Disk_1,
		2:  disks.Disk_2,
		3:  disks.Disk_3,
		4:  disks.Disk_4,
		5:  disks.Disk_5,
		6:  disks.Disk_6,
		7:  disks.Disk_7,
		8:  disks.Disk_8,
		9:  disks.Disk_9,
		10: disks.Disk_10,
		11: disks.Disk_11,
		12: disks.Disk_12,
		13: disks.Disk_13,
		14: disks.Disk_14,
		15: disks.Disk_15,
	}
}

// TODO write test
func (QemuVirtIODisks) mapToStruct(params map[string]interface{}) *QemuVirtIODisks {
	disks := QemuVirtIODisks{}
	var structPopulated bool
	if _, isSet := params["virtio0"]; isSet {
		disks.Disk_0 = QemuVirtIOStorage{}.mapToStruct(params["virtio0"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio1"]; isSet {
		disks.Disk_1 = QemuVirtIOStorage{}.mapToStruct(params["virtio1"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio2"]; isSet {
		disks.Disk_2 = QemuVirtIOStorage{}.mapToStruct(params["virtio2"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio3"]; isSet {
		disks.Disk_3 = QemuVirtIOStorage{}.mapToStruct(params["virtio3"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio4"]; isSet {
		disks.Disk_4 = QemuVirtIOStorage{}.mapToStruct(params["virtio4"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio5"]; isSet {
		disks.Disk_5 = QemuVirtIOStorage{}.mapToStruct(params["virtio5"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio6"]; isSet {
		disks.Disk_6 = QemuVirtIOStorage{}.mapToStruct(params["virtio6"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio7"]; isSet {
		disks.Disk_7 = QemuVirtIOStorage{}.mapToStruct(params["virtio7"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio8"]; isSet {
		disks.Disk_8 = QemuVirtIOStorage{}.mapToStruct(params["virtio8"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio9"]; isSet {
		disks.Disk_9 = QemuVirtIOStorage{}.mapToStruct(params["virtio9"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio10"]; isSet {
		disks.Disk_10 = QemuVirtIOStorage{}.mapToStruct(params["virtio10"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio11"]; isSet {
		disks.Disk_11 = QemuVirtIOStorage{}.mapToStruct(params["virtio11"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio12"]; isSet {
		disks.Disk_12 = QemuVirtIOStorage{}.mapToStruct(params["virtio12"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio13"]; isSet {
		disks.Disk_13 = QemuVirtIOStorage{}.mapToStruct(params["virtio13"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio14"]; isSet {
		disks.Disk_14 = QemuVirtIOStorage{}.mapToStruct(params["virtio14"].(string))
		structPopulated = true
	}
	if _, isSet := params["virtio15"]; isSet {
		disks.Disk_15 = QemuVirtIOStorage{}.mapToStruct(params["virtio15"].(string))
		structPopulated = true
	}
	if structPopulated {
		return &disks
	}
	return nil
}

func (disks QemuVirtIODisks) Validate() error {
	diskMap := disks.mapToIntMap()
	for _, e := range diskMap {
		if e != nil {
			err := e.Validate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type QemuVirtIOPassthrough struct {
	AsyncIO   QemuDiskAsyncIO
	Backup    bool
	Bandwidth QemuDiskBandwidth
	Cache     QemuDiskCache
	Discard   bool
	File      string
	IOThread  bool
	ReadOnly  bool
	Serial    QemuDiskSerial `json:"serial,omitempty"`
	Size      uint           //size is only returned and setting it has no effect
}

func (passthrough *QemuVirtIOPassthrough) convertDataStructure() *qemuDisk {
	return &qemuDisk{
		AsyncIO:   passthrough.AsyncIO,
		Backup:    passthrough.Backup,
		Bandwidth: passthrough.Bandwidth,
		Cache:     passthrough.Cache,
		Discard:   passthrough.Discard,
		File:      passthrough.File,
		IOThread:  passthrough.IOThread,
		ReadOnly:  passthrough.ReadOnly,
		Serial:    passthrough.Serial,
		Type:      virtIO,
	}
}

func (passthrough QemuVirtIOPassthrough) Validate() error {
	return passthrough.convertDataStructure().validate()
}

type QemuVirtIOStorage struct {
	CdRom       *QemuCdRom
	CloudInit   *QemuCloudInitDisk
	Disk        *QemuVirtIODisk
	Passthrough *QemuVirtIOPassthrough
}

// TODO write test
// converts to qemuStorage
func (storage *QemuVirtIOStorage) convertDataStructure() *qemuStorage {
	if storage == nil {
		return nil
	}
	generalizedStorage := qemuStorage{
		CdRom:     storage.CdRom,
		CloudInit: storage.CloudInit,
	}
	if storage.Disk != nil {
		generalizedStorage.Disk = storage.Disk.convertDataStructure()
	}
	if storage.Passthrough != nil {
		generalizedStorage.Passthrough = storage.Passthrough.convertDataStructure()
	}
	return &generalizedStorage
}

// TODO write test
func (QemuVirtIOStorage) mapToStruct(param string) *QemuVirtIOStorage {
	settings := splitStringOfSettings(param)
	tmpCdRom := qemuCdRom{}.mapToStruct(settings)
	if tmpCdRom != nil {
		if tmpCdRom.CdRom {
			return &QemuVirtIOStorage{CdRom: QemuCdRom{}.mapToStruct(*tmpCdRom)}
		} else {
			return &QemuVirtIOStorage{CloudInit: QemuCloudInitDisk{}.mapToStruct(*tmpCdRom)}
		}
	}

	tmpDisk := qemuDisk{}.mapToStruct(settings)
	if tmpDisk == nil {
		return nil
	}
	if tmpDisk.File == "" {
		return &QemuVirtIOStorage{Disk: &QemuVirtIODisk{
			AsyncIO:   tmpDisk.AsyncIO,
			Backup:    tmpDisk.Backup,
			Bandwidth: tmpDisk.Bandwidth,
			Cache:     tmpDisk.Cache,
			Discard:   tmpDisk.Discard,
			Format:    tmpDisk.Format,
			Id:        tmpDisk.Id,
			IOThread:  tmpDisk.IOThread,
			ReadOnly:  tmpDisk.ReadOnly,
			Replicate: tmpDisk.Replicate,
			Serial:    tmpDisk.Serial,
			Size:      tmpDisk.Size,
			Storage:   tmpDisk.Storage,
		}}
	}
	return &QemuVirtIOStorage{Passthrough: &QemuVirtIOPassthrough{
		AsyncIO:   tmpDisk.AsyncIO,
		Backup:    tmpDisk.Backup,
		Bandwidth: tmpDisk.Bandwidth,
		Cache:     tmpDisk.Cache,
		Discard:   tmpDisk.Discard,
		File:      tmpDisk.File,
		IOThread:  tmpDisk.IOThread,
		ReadOnly:  tmpDisk.ReadOnly,
		Serial:    tmpDisk.Serial,
		Size:      tmpDisk.Size,
	}}
}

func (storage QemuVirtIOStorage) Validate() (err error) {
	// First check if more than one item is nil
	var subTypeSet bool
	if storage.CdRom != nil {
		subTypeSet = true
	}
	if storage.CloudInit != nil {
		if err = diskSubtypeSet(subTypeSet); err != nil {
			return
		}
		subTypeSet = true
	}
	if storage.Disk != nil {
		if err = diskSubtypeSet(subTypeSet); err != nil {
			return
		}
		subTypeSet = true
	}
	if storage.Passthrough != nil {
		if err = diskSubtypeSet(subTypeSet); err != nil {
			return
		}
	}
	// Validate sub items
	if storage.CdRom != nil {
		if err = storage.CdRom.Validate(); err != nil {
			return
		}
	}
	if storage.CloudInit != nil {
		if err = storage.CloudInit.Validate(); err != nil {
			return
		}
	}
	if storage.Disk != nil {
		if err = storage.Disk.Validate(); err != nil {
			return
		}
	}
	if storage.Passthrough != nil {
		err = storage.Passthrough.Validate()
	}
	return
}
