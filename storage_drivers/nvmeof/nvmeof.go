package nvmeof

import (
    "fmt"
    "encoding/json"
    "os/exec"
	"context"
	"strconv"
	"strings"
    
	tridentconfig "github.com/netapp/trident/config"
    "github.com/netapp/trident/utils"
	. "github.com/netapp/trident/logging"
	sa "github.com/netapp/trident/storage_attribute"
	drivers "github.com/netapp/trident/storage_drivers"
	storage "github.com/netapp/trident/storage"

	"github.com/RoaringBitmap/roaring"
)


const (
	StorageDriverName = "nvmeof"
)

func (d *NVMeoFBackend) Name() string {
	return StorageDriverName
}

type NVMeoFBackend struct {
    config NVMeoFConfig
    volumes map[string]*storage.VolumeConfig
}

type Namespace struct {
    Name string
    Size string
    NamespaceID string

}

func (d *NVMeoFBackend) BackendName() string {
    return "nvmeof"
}

func NewNVMeoFBackend(config NVMeoFConfig) (*NVMeoFBackend, error) {
	cmd := exec.Command("nvme", "discover", "-t", config.TransportType, "-a", config.TargetAddress, "-s", config.TargetPort)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return nil, fmt.Errorf("NVMe-oF discovery failed: %v, output: %s", err, output)
    }
    return &NVMeoFBackend{config: config}, nil
}

func (d *NVMeoFBackend) Create(ctx context.Context, volConfig *storage.VolumeConfig, storagePool storage.Pool, volAttributes map[string]sa.Request) error{
    cmd := exec.Command("nvme", "create-ns", "--size", fmt.Sprintf("%d", volConfig.Size))
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error creating volume: %v, output: %s", err, output)
    }
    return nil
}

func (d *NVMeoFBackend) Initialize(
    ctx context.Context, driverContext tridentconfig.DriverContext, configJSON string,
    commonConfig *drivers.CommonStorageDriverConfig, backendSecret map[string]string, backendUUID string,
) error {
    if err := json.Unmarshal([]byte(configJSON), &d.config); err != nil {
        return fmt.Errorf("unable to parse NVMe-oF configuration: %v", err)
    }

    if err := d.connectToNVMeTarget(); err != nil {
        return fmt.Errorf("unable to connect to NVMe-oF target: %v", err)
    }
    d.initializeVolumeTracking(ctx)

    //d.backendUUID = backendUUID

    Logc(ctx).WithFields(LogFields{
        "driverName": commonConfig.StorageDriverName,
        "backendUUID": backendUUID,
    }).Info("Storage driver initialized")

    return nil
}

func (d *NVMeoFBackend) connectToNVMeTarget() error {
    cmd := exec.Command("nvme", "connect", "-t", d.config.TransportType, "-a", d.config.TargetAddress, "-s", d.config.TargetPort)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error connecting to NVMe-oF target: %v, output: %s", err, string(output))
    }
    return nil
}

func (d *NVMeoFBackend) initializeVolumeTracking(ctx context.Context) {
	if d.volumes == nil {
        d.volumes = make(map[string]*storage.VolumeConfig)
    }

    cmd := exec.Command("nvme", "list")
    output, err := cmd.CombinedOutput()
    if err != nil {
        Logc(ctx).Errorf("Failed to list NVMe namespaces: %v, output: %s", err, string(output))
        return
    }
    namespaces := parseNVMeListOutput(output)

    for _, ns := range namespaces {
        vol := &storage.VolumeConfig{
            Name:      ns.Name,
            Size:      ns.Size,
            InternalID: ns.NamespaceID,
        }
        d.volumes[ns.Name] = vol
    }

    Logc(ctx).Info("Successfully initialized volume tracking with %d volumes.", len(d.volumes))

}

func parseNVMeListOutput(output []byte) []Namespace {
	var namespaces []Namespace
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		fields := strings.Fields(line)

		
		if len(fields) >= 3 {
			size, err := strconv.ParseInt(fields[2], 10, 64)
			if err != nil {
				continue
			}
			namespace := Namespace{
				NamespaceID: fields[0],
				Name:        fields[1],
				Size:        strconv.FormatInt(size, 10),
			}
			namespaces = append(namespaces, namespace)
		}
	}

	return namespaces
}

func (backend *NVMeoFBackend) Destroy(
    ctx context.Context, 
    volConfig *storage.VolumeConfig) error {
    cmd := exec.Command("nvme", "delete-ns", volConfig.Name)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error deleting volume: %v, output: %s", err, output)
    }
    return nil
}

func (backend *NVMeoFBackend) Resize(ctx context.Context, volConfig *storage.VolumeConfig, sizeBytes uint64) error {
    cmd := exec.Command("nvme", "resize", volConfig.Name, fmt.Sprintf("--size=%d", volConfig.Size))
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error resizing volume: %v, output: %s", err, output)
    }
    return nil
}

func (backend *NVMeoFBackend) AttachVolume(volumeName string, hostName string) error {
    cmd := exec.Command("nvme", "connect", "--volume", volumeName, "--host", hostName)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error attaching volume to host: %v, output: %s", err, output)
    }
    return nil
}

func (backend *NVMeoFBackend) DetachVolume(volumeName string) error {
    cmd := exec.Command("nvme", "disconnect", "--volume", volumeName)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error detaching volume: %v, output: %s", err, output)
    }
    return nil
}

func (d *NVMeoFBackend) init() {
    //drivers.RegisterDriver("nvmeof", NewNVMeoFBackend)
}

func (d *NVMeoFBackend) Terminate(ctx context.Context, backendUUID string){}


func (d *NVMeoFBackend) CreatePrepare(ctx context.Context, volConfig *storage.VolumeConfig){

}
func (d *NVMeoFBackend) CreateFollowup(ctx context.Context, volConfig *storage.VolumeConfig) error{
    return nil
}
func (d *NVMeoFBackend) CreateClone(ctx context.Context, sourceVolConfig, cloneVolConfig *storage.VolumeConfig, storagePool storage.Pool) error{
    return nil
}
func (d *NVMeoFBackend) Import(ctx context.Context, volConfig *storage.VolumeConfig, originalName string) error{
    return nil
}
func (d *NVMeoFBackend) Rename(ctx context.Context, name, newName string) error{
    return nil
}
func (d *NVMeoFBackend) Get(ctx context.Context, name string) error{
    return nil
}
func (d *NVMeoFBackend) GetInternalVolumeName(ctx context.Context, name string) string{
    return ""
}
func (d *NVMeoFBackend) GetStorageBackendSpecs(ctx context.Context, backend storage.Backend) error{
    return nil
}
func (d *NVMeoFBackend) GetStorageBackendPhysicalPoolNames(ctx context.Context) []string{
    return nil
}
func (d *NVMeoFBackend) GetProtocol(ctx context.Context) tridentconfig.Protocol{
    return tridentconfig.Block
}
func (d *NVMeoFBackend) Publish(ctx context.Context, volConfig *storage.VolumeConfig, publishInfo *utils.VolumePublishInfo) error{
    return nil
}
func (d *NVMeoFBackend) CanSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) error{
    return nil
}
func (d *NVMeoFBackend) GetSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) (*storage.Snapshot, error){
    return nil, nil
}
func (d *NVMeoFBackend) GetSnapshots(ctx context.Context, volConfig *storage.VolumeConfig) ([]*storage.Snapshot, error){
    return nil, nil
}
func (d *NVMeoFBackend) CreateSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) (*storage.Snapshot, error){
    return nil, nil
}
func (d *NVMeoFBackend) RestoreSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) error{
    return nil
}
func (d *NVMeoFBackend) DeleteSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) error{
    return nil
}
func (d *NVMeoFBackend) StoreConfig(ctx context.Context, b *storage.PersistentStorageBackendConfig){
}
func (d *NVMeoFBackend) GetExternalConfig(ctx context.Context) interface{}{
    return d.config
}
func (d *NVMeoFBackend) GetVolumeExternal(ctx context.Context, name string) (*storage.VolumeExternal, error){
    return nil, nil
}
func (d *NVMeoFBackend) GetVolumeExternalWrappers(context.Context, chan *storage.VolumeExternalWrapper){
}
func (d *NVMeoFBackend) GetUpdateType(ctx context.Context, driver storage.Driver) *roaring.Bitmap{
    return nil
}
func (d *NVMeoFBackend) ReconcileNodeAccess(ctx context.Context, nodes []*utils.Node, backendUUID, tridentUUID string) error{
    return nil
}
func (d *NVMeoFBackend) GetCommonConfig(context.Context) *drivers.CommonStorageDriverConfig{
    return nil
}

func (d *NVMeoFBackend) Initialized() bool{
    return false
}
