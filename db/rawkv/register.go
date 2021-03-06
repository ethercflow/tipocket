package rawkv

//
//var (
//	register = []byte("acc")
//)
//
//type registerClient struct {
//	db *tikv.RawKVClient
//	r  *rand.Rand
//}
//
//func (c *registerClient) SetUp(ctx context.Context, nodes []clusterTypes.Node, node clusterTypes.Node) error {
//	c.r = rand.New(rand.NewSource(time.Now().UnixNano()))
//	db, err := tikv.NewRawKVClient([]string{fmt.Sprintf("%s:2379", node.IP)}, config.Security{})
//	if err != nil {
//		return err
//	}
//
//	c.db = db
//
//	// Do SetUp in the first node
//	if node != nodes[0] {
//		return nil
//	}
//
//	log.Printf("begin to initial register on node %s", node.IP)
//
//	db.Put(register, []byte("0"))
//
//	return nil
//}
//
//func (c *registerClient) TearDown(ctx context.Context, nodes []clusterTypes.Node, node clusterTypes.Node) error {
//	return c.db.Close()
//}
//
//func (c *registerClient) invokeRead(ctx context.Context, r model.RegisterRequest) model.RegisterResponse {
//	val, err := c.db.Get(register)
//	if err != nil {
//		return model.RegisterResponse{Unknown: true}
//	}
//	v, err := strconv.ParseInt(string(val), 10, 64)
//	if err != nil {
//		panic(fmt.Sprintf("invalid value: %s", val))
//	}
//	return model.RegisterResponse{Value: int(v)}
//}
//
//func (c *registerClient) Invoke(ctx context.Context, node clusterTypes.Node, r interface{}) interface{} {
//	arg := r.(model.RegisterRequest)
//	if arg.Op == model.RegisterRead {
//		return c.invokeRead(ctx, arg)
//	}
//
//	val := fmt.Sprintf("%d", arg.Value)
//	err := c.db.Put(register, []byte(val))
//	if err != nil {
//		return model.RegisterResponse{Unknown: true}
//	}
//	return model.RegisterResponse{}
//}
//
//func (c *registerClient) NextRequest() interface{} {
//	r := model.RegisterRequest{
//		Op: c.r.Intn(2) == 1,
//	}
//	if r.Op == model.RegisterRead {
//		return r
//	}
//
//	r.Value = int(c.r.Int63())
//	return r
//}
//
//// DumpState the database state(also the model's state)
//func (c *registerClient) DumpState(ctx context.Context) (interface{}, error) {
//	val, err := c.db.Get(register)
//	if err != nil {
//		return nil, err
//	}
//	v, err := strconv.ParseInt(string(val), 10, 64)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func newRegisterEvent(v interface{}, id uint) porcupine.Event {
//	if _, ok := v.(model.RegisterRequest); ok {
//		return porcupine.Event{Kind: porcupine.CallEvent, Value: v, Id: id}
//	}
//
//	return porcupine.Event{Kind: porcupine.ReturnEvent, Value: v, Id: id}
//}
//
//// RegisterClientCreator creates a register test client for rawkv.
//type RegisterClientCreator struct {
//}
//
//// Create creates a client.
//func (RegisterClientCreator) Create(node clusterTypes.Node) core.Client {
//	return &registerClient{}
//}
