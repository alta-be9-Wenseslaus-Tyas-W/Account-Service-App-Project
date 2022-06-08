package entities

type Transfers struct {
	IdTransfers    int
	IdUserPengirim int
	IdUserPenerima int
	Nominal        int
	SisaSaldo      int
}

type HistoryTrasfer struct {
	NamaPengirim string
	NamaPenerima string
	Nominal      int
	SisaSaldo    int
}
