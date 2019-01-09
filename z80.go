package main

// Clock Any clock
type clock struct {
	m uint16
	t uint16
}

// Reg The register
type reg struct {
	a, b, c, d, e, h, l, f, r uint8
	pc, sp                    uint16
	lc                        clock
}

type opFunc func(*Z80, *MMU)

// Op An operation
type op struct {
	f opFunc
	c clock
}

func (op *op) exec(z80 *Z80, mmu *MMU) {
	op.f(z80, mmu)
	z80.r.lc.m = op.c.m
	z80.r.lc.t = op.c.t
}

// Z80 The Z80
type Z80 struct {
	r      reg
	c      clock
	ops    [0x100]interface{}
	clocks [0x100]clock
}

// Init initializes the Z80
func (z80 *Z80) Init() {
	// No code
	z80.ops = [0x100]interface{}{
		opsNOP,
		opsLDBCnn,
		opsLDBCmA,
		opsINCBC,
		opsINCrb,
		opsDECrb,
		opsLDrnb,
		opsRLCA,
		opsLDmmSP,
		opsADDHLBC,
		opsLDABCm,
		opsDECBC,
		opsINCrc,
		opsDECrc,
		opsLDrnc,
		opsRRCA,

		// 10
		opsDJNZn,
		opsLDDEnn,
		opsLDDEmA,
		opsINCDE,
		opsINCrd,
		opsDECrd,
		opsLDrnd,
		opsRLA,
		opsJRn,
		opsADDHLDE,
		opsLDADEm,
		opsDECDE,
		opsINCre,
		opsDECre,
		opsLDrne,
		opsRRA,

		// 20
		opsJRNZn,
		opsLDHLnn,
		opsLDHLIA,
		opsINCHL,
		opsINCrh,
		opsDECrh,
		opsLDrnh,
		opsXX,
		opsJRZn,
		opsADDHLHL,
		opsLDAHLI,
		opsDECHL,
		opsINCrl,
		opsDECrl,
		opsLDrnl,
		opsCPL,

		// 30
		opsJRNCn,
		opsLDSPnn,
		opsLDHLDA,
		opsINCSP,
		opsINCHLm,
		opsDECHLm,
		opsLDHLmn,
		opsSCF,
		opsJRCn,
		opsADDHLSP,
		opsLDAHLD,
		opsDECSP,
		opsINCra,
		opsDECra,
		opsLDrna,
		opsCCF,

		// 40
		opsLDrrbb,
		opsLDrrbc,
		opsLDrrbd,
		opsLDrrbe,
		opsLDrrbh,
		opsLDrrbl,
		opsLDrHLmb,
		opsLDrrba,
		opsLDrrcb,
		opsLDrrcc,
		opsLDrrcd,
		opsLDrrce,
		opsLDrrch,
		opsLDrrcl,
		opsLDrHLmc,
		opsLDrrca,

		// 50
		opsLDrrdb,
		opsLDrrdc,
		opsLDrrdd,
		opsLDrrde,
		opsLDrrdh,
		opsLDrrdl,
		opsLDrHLmd,
		opsLDrrda,
		opsLDrreb,
		opsLDrrec,
		opsLDrred,
		opsLDrree,
		opsLDrreh,
		opsLDrrel,
		opsLDrHLme,
		opsLDrrea,

		// 60
		opsLDrrhb,
		opsLDrrhc,
		opsLDrrhd,
		opsLDrrhe,
		opsLDrrhh,
		opsLDrrhl,
		opsLDrHLmh,
		opsLDrrha,
		opsLDrrlb,
		opsLDrrlc,
		opsLDrrld,
		opsLDrrle,
		opsLDrrlh,
		opsLDrrll,
		opsLDrHLml,
		opsLDrrla,

		// 70
		opsLDHLmrb,
		opsLDHLmrc,
		opsLDHLmrd,
		opsLDHLmre,
		opsLDHLmrh,
		opsLDHLmrl,
		opsHALT,
		opsLDHLmra,
		opsLDrrab,
		opsLDrrac,
		opsLDrrad,
		opsLDrrae,
		opsLDrrah,
		opsLDrral,
		opsLDrHLma,
		opsLDrraa,

		// 80
		opsADDrb,
		opsADDrc,
		opsADDrd,
		opsADDre,
		opsADDrh,
		opsADDrl,
		opsADDHL,
		opsADDra,
		opsADCrb,
		opsADCrc,
		opsADCrd,
		opsADCre,
		opsADCrh,
		opsADCrl,
		opsADCHL,
		opsADCra,

		// 90
		opsSUBrb,
		opsSUBrc,
		opsSUBrd,
		opsSUBre,
		opsSUBrh,
		opsSUBrl,
		opsSUBHL,
		opsSUBra,
		opsSBCrb,
		opsSBCrc,
		opsSBCrd,
		opsSBCre,
		opsSBCrh,
		opsSBCrl,
		opsSBCHL,
		opsSBCra,

		// A0
		opsANDrb,
		opsANDrc,
		opsANDrd,
		opsANDre,
		opsANDrh,
		opsANDrl,
		opsANDHL,
		opsANDra,
		opsXORrb,
		opsXORrc,
		opsXORrd,
		opsXORre,
		opsXORrh,
		opsXORrl,
		opsXORHL,
		opsXORra,

		// B0
		opsORrb,
		opsORrc,
		opsORrd,
		opsORre,
		opsORrh,
		opsORrl,
		opsORHL,
		opsORra,
		opsCPrb,
		opsCPrc,
		opsCPrd,
		opsCPre,
		opsCPrh,
		opsCPrl,
		opsCPHL,
		opsCPra,

		// C0
		opsRETNZ,
		opsPOPBC,
		opsJPNZnn,
		opsJPnn,
		opsCALLNZnn,
		opsPUSHBC,
		opsADDn,
		opsRST00,
		opsRETZ,
		opsRET,
		opsJPZnn,
		opsMAPcb,
		opsCALLZnn,
		opsCALLnn,
		opsADCn,
		opsRST08,

		// D0
		opsRETNC,
		opsPOPDE,
		opsJPNCnn,
		opsXX,
		opsCALLNCnn,
		opsPUSHDE,
		opsSUBn,
		opsRST10,
		opsRETC,
		opsRETI,
		opsJPCnn,
		opsXX,
		opsCALLCnn,
		opsXX,
		opsSBCn,
		opsRST18,

		// E0
		opsLDIOnA,
		opsPOPHL,
		opsLDIOCA,
		opsXX,
		opsXX,
		opsPUSHHL,
		opsANDn,
		opsRST20,
		opsADDSPn,
		opsJPHL,
		opsLDmmA,
		opsXX,
		opsXX,
		opsXX,
		opsORn,
		opsRST28,

		// F0
		opsLDAIOn,
		opsPOPAF,
		opsLDAIOC,
		opsDI,
		opsXX,
		opsPUSHAF,
		opsXORn,
		opsRST30,
		opsLDHLSPn,
		opsXX,
		opsLDAmm,
		opsEI,
		opsXX,
		opsXX,
		opsCPn,
		opsRST38}
}

// Exec executes the op
func (z80 *Z80) Exec(op uint8, mmu *MMU) {
	// NOP
	f := z80.ops[op]
	f.(func(*Z80, *MMU))(z80, mmu)
	c := z80.clocks[op]
	z80.r.lc = c
}

// Ops

func opsNOP(z80 *Z80, mmu *MMU) {
	// NOP
}

func opsLDBCnn(z80 *Z80, mmu *MMU) {
	z80.r.c = mmu.rb(z80.r.pc)
	z80.r.b = mmu.rb(z80.r.pc + 1)
	z80.r.pc += 2
}

func opsLDBCmA(z80 *Z80, mmu *MMU) {
	loc := uint16(z80.r.b)<<8 + uint16(z80.r.c)
	mmu.wb(loc, z80.r.a)
}

func opsINCBC(z80 *Z80, mmu *MMU) {

}

func opsINCrb(z80 *Z80, mmu *MMU) {

}

func opsDECrb(z80 *Z80, mmu *MMU) {

}

func opsLDrnb(z80 *Z80, mmu *MMU) {

}

func opsRLCA(z80 *Z80, mmu *MMU) {

}

func opsLDmmSP(z80 *Z80, mmu *MMU) {

}

func opsADDHLBC(z80 *Z80, mmu *MMU) {

}

func opsLDABCm(z80 *Z80, mmu *MMU) {

}

func opsDECBC(z80 *Z80, mmu *MMU) {

}

func opsINCrc(z80 *Z80, mmu *MMU) {

}

func opsDECrc(z80 *Z80, mmu *MMU) {

}

func opsLDrnc(z80 *Z80, mmu *MMU) {

}

func opsRRCA(z80 *Z80, mmu *MMU) {

}

// 10
func opsDJNZn(z80 *Z80, mmu *MMU) {

}

func opsLDDEnn(z80 *Z80, mmu *MMU) {

}

func opsLDDEmA(z80 *Z80, mmu *MMU) {

}

func opsINCDE(z80 *Z80, mmu *MMU) {

}

func opsINCrd(z80 *Z80, mmu *MMU) {

}

func opsDECrd(z80 *Z80, mmu *MMU) {

}

func opsLDrnd(z80 *Z80, mmu *MMU) {

}

func opsRLA(z80 *Z80, mmu *MMU) {

}

func opsJRn(z80 *Z80, mmu *MMU) {

}

func opsADDHLDE(z80 *Z80, mmu *MMU) {

}

func opsLDADEm(z80 *Z80, mmu *MMU) {

}

func opsDECDE(z80 *Z80, mmu *MMU) {

}

func opsINCre(z80 *Z80, mmu *MMU) {

}

func opsDECre(z80 *Z80, mmu *MMU) {

}

func opsLDrne(z80 *Z80, mmu *MMU) {

}

func opsRRA(z80 *Z80, mmu *MMU) {

}

// 20
func opsJRNZn(z80 *Z80, mmu *MMU) {

}

func opsLDHLnn(z80 *Z80, mmu *MMU) {

}

func opsLDHLIA(z80 *Z80, mmu *MMU) {

}

func opsINCHL(z80 *Z80, mmu *MMU) {

}

func opsINCrh(z80 *Z80, mmu *MMU) {

}

func opsDECrh(z80 *Z80, mmu *MMU) {

}

func opsLDrnh(z80 *Z80, mmu *MMU) {

}

func opsXX(z80 *Z80, mmu *MMU) {

}

func opsJRZn(z80 *Z80, mmu *MMU) {

}

func opsADDHLHL(z80 *Z80, mmu *MMU) {

}

func opsLDAHLI(z80 *Z80, mmu *MMU) {

}

func opsDECHL(z80 *Z80, mmu *MMU) {

}

func opsINCrl(z80 *Z80, mmu *MMU) {

}

func opsDECrl(z80 *Z80, mmu *MMU) {

}

func opsLDrnl(z80 *Z80, mmu *MMU) {

}

func opsCPL(z80 *Z80, mmu *MMU) {

}

// 30
func opsJRNCn(z80 *Z80, mmu *MMU) {

}

func opsLDSPnn(z80 *Z80, mmu *MMU) {

}

func opsLDHLDA(z80 *Z80, mmu *MMU) {

}

func opsINCSP(z80 *Z80, mmu *MMU) {

}

func opsINCHLm(z80 *Z80, mmu *MMU) {

}

func opsDECHLm(z80 *Z80, mmu *MMU) {

}

func opsLDHLmn(z80 *Z80, mmu *MMU) {

}

func opsSCF(z80 *Z80, mmu *MMU) {

}

func opsJRCn(z80 *Z80, mmu *MMU) {

}

func opsADDHLSP(z80 *Z80, mmu *MMU) {

}

func opsLDAHLD(z80 *Z80, mmu *MMU) {

}

func opsDECSP(z80 *Z80, mmu *MMU) {

}

func opsINCra(z80 *Z80, mmu *MMU) {

}

func opsDECra(z80 *Z80, mmu *MMU) {

}

func opsLDrna(z80 *Z80, mmu *MMU) {

}

func opsCCF(z80 *Z80, mmu *MMU) {

}

// 40
func opsLDrrbb(z80 *Z80, mmu *MMU) {

}

func opsLDrrbc(z80 *Z80, mmu *MMU) {

}

func opsLDrrbd(z80 *Z80, mmu *MMU) {

}

func opsLDrrbe(z80 *Z80, mmu *MMU) {

}

func opsLDrrbh(z80 *Z80, mmu *MMU) {

}

func opsLDrrbl(z80 *Z80, mmu *MMU) {

}

func opsLDrHLmb(z80 *Z80, mmu *MMU) {

}

func opsLDrrba(z80 *Z80, mmu *MMU) {

}

func opsLDrrcb(z80 *Z80, mmu *MMU) {

}

func opsLDrrcc(z80 *Z80, mmu *MMU) {

}

func opsLDrrcd(z80 *Z80, mmu *MMU) {

}

func opsLDrrce(z80 *Z80, mmu *MMU) {

}

func opsLDrrch(z80 *Z80, mmu *MMU) {

}

func opsLDrrcl(z80 *Z80, mmu *MMU) {

}

func opsLDrHLmc(z80 *Z80, mmu *MMU) {

}

func opsLDrrca(z80 *Z80, mmu *MMU) {

}

// 50
func opsLDrrdb(z80 *Z80, mmu *MMU) {

}

func opsLDrrdc(z80 *Z80, mmu *MMU) {

}

func opsLDrrdd(z80 *Z80, mmu *MMU) {

}

func opsLDrrde(z80 *Z80, mmu *MMU) {

}

func opsLDrrdh(z80 *Z80, mmu *MMU) {

}

func opsLDrrdl(z80 *Z80, mmu *MMU) {

}

func opsLDrHLmd(z80 *Z80, mmu *MMU) {

}

func opsLDrrda(z80 *Z80, mmu *MMU) {

}

func opsLDrreb(z80 *Z80, mmu *MMU) {

}

func opsLDrrec(z80 *Z80, mmu *MMU) {

}

func opsLDrred(z80 *Z80, mmu *MMU) {

}

func opsLDrree(z80 *Z80, mmu *MMU) {

}

func opsLDrreh(z80 *Z80, mmu *MMU) {

}

func opsLDrrel(z80 *Z80, mmu *MMU) {

}

func opsLDrHLme(z80 *Z80, mmu *MMU) {

}

func opsLDrrea(z80 *Z80, mmu *MMU) {

}

// 60
func opsLDrrhb(z80 *Z80, mmu *MMU) {

}

func opsLDrrhc(z80 *Z80, mmu *MMU) {

}

func opsLDrrhd(z80 *Z80, mmu *MMU) {

}

func opsLDrrhe(z80 *Z80, mmu *MMU) {

}

func opsLDrrhh(z80 *Z80, mmu *MMU) {

}

func opsLDrrhl(z80 *Z80, mmu *MMU) {

}

func opsLDrHLmh(z80 *Z80, mmu *MMU) {

}

func opsLDrrha(z80 *Z80, mmu *MMU) {

}

func opsLDrrlb(z80 *Z80, mmu *MMU) {

}

func opsLDrrlc(z80 *Z80, mmu *MMU) {

}

func opsLDrrld(z80 *Z80, mmu *MMU) {

}

func opsLDrrle(z80 *Z80, mmu *MMU) {

}

func opsLDrrlh(z80 *Z80, mmu *MMU) {

}

func opsLDrrll(z80 *Z80, mmu *MMU) {

}

func opsLDrHLml(z80 *Z80, mmu *MMU) {

}

func opsLDrrla(z80 *Z80, mmu *MMU) {

}

// 70
func opsLDHLmrb(z80 *Z80, mmu *MMU) {

}

func opsLDHLmrc(z80 *Z80, mmu *MMU) {

}

func opsLDHLmrd(z80 *Z80, mmu *MMU) {

}

func opsLDHLmre(z80 *Z80, mmu *MMU) {

}

func opsLDHLmrh(z80 *Z80, mmu *MMU) {

}

func opsLDHLmrl(z80 *Z80, mmu *MMU) {

}

func opsHALT(z80 *Z80, mmu *MMU) {

}

func opsLDHLmra(z80 *Z80, mmu *MMU) {

}

func opsLDrrab(z80 *Z80, mmu *MMU) {

}

func opsLDrrac(z80 *Z80, mmu *MMU) {

}

func opsLDrrad(z80 *Z80, mmu *MMU) {

}

func opsLDrrae(z80 *Z80, mmu *MMU) {

}

func opsLDrrah(z80 *Z80, mmu *MMU) {

}

func opsLDrral(z80 *Z80, mmu *MMU) {

}

func opsLDrHLma(z80 *Z80, mmu *MMU) {

}

func opsLDrraa(z80 *Z80, mmu *MMU) {

}

// 80
func opsADDrb(z80 *Z80, mmu *MMU) {

}

func opsADDrc(z80 *Z80, mmu *MMU) {

}

func opsADDrd(z80 *Z80, mmu *MMU) {

}

func opsADDre(z80 *Z80, mmu *MMU) {

}

func opsADDrh(z80 *Z80, mmu *MMU) {

}

func opsADDrl(z80 *Z80, mmu *MMU) {

}

func opsADDHL(z80 *Z80, mmu *MMU) {

}

func opsADDra(z80 *Z80, mmu *MMU) {

}

func opsADCrb(z80 *Z80, mmu *MMU) {

}

func opsADCrc(z80 *Z80, mmu *MMU) {

}

func opsADCrd(z80 *Z80, mmu *MMU) {

}

func opsADCre(z80 *Z80, mmu *MMU) {

}

func opsADCrh(z80 *Z80, mmu *MMU) {

}

func opsADCrl(z80 *Z80, mmu *MMU) {

}

func opsADCHL(z80 *Z80, mmu *MMU) {

}

func opsADCra(z80 *Z80, mmu *MMU) {

}

// 90
func opsSUBrb(z80 *Z80, mmu *MMU) {

}

func opsSUBrc(z80 *Z80, mmu *MMU) {

}

func opsSUBrd(z80 *Z80, mmu *MMU) {

}

func opsSUBre(z80 *Z80, mmu *MMU) {

}

func opsSUBrh(z80 *Z80, mmu *MMU) {

}

func opsSUBrl(z80 *Z80, mmu *MMU) {

}

func opsSUBHL(z80 *Z80, mmu *MMU) {

}

func opsSUBra(z80 *Z80, mmu *MMU) {

}

func opsSBCrb(z80 *Z80, mmu *MMU) {

}

func opsSBCrc(z80 *Z80, mmu *MMU) {

}

func opsSBCrd(z80 *Z80, mmu *MMU) {

}

func opsSBCre(z80 *Z80, mmu *MMU) {

}

func opsSBCrh(z80 *Z80, mmu *MMU) {

}

func opsSBCrl(z80 *Z80, mmu *MMU) {

}

func opsSBCHL(z80 *Z80, mmu *MMU) {

}

func opsSBCra(z80 *Z80, mmu *MMU) {

}

// A0
func opsANDrb(z80 *Z80, mmu *MMU) {

}

func opsANDrc(z80 *Z80, mmu *MMU) {

}

func opsANDrd(z80 *Z80, mmu *MMU) {

}

func opsANDre(z80 *Z80, mmu *MMU) {

}

func opsANDrh(z80 *Z80, mmu *MMU) {

}

func opsANDrl(z80 *Z80, mmu *MMU) {

}

func opsANDHL(z80 *Z80, mmu *MMU) {

}

func opsANDra(z80 *Z80, mmu *MMU) {

}

func opsXORrb(z80 *Z80, mmu *MMU) {

}

func opsXORrc(z80 *Z80, mmu *MMU) {

}

func opsXORrd(z80 *Z80, mmu *MMU) {

}

func opsXORre(z80 *Z80, mmu *MMU) {

}

func opsXORrh(z80 *Z80, mmu *MMU) {

}

func opsXORrl(z80 *Z80, mmu *MMU) {

}

func opsXORHL(z80 *Z80, mmu *MMU) {

}

func opsXORra(z80 *Z80, mmu *MMU) {

}

// B0
func opsORrb(z80 *Z80, mmu *MMU) {

}

func opsORrc(z80 *Z80, mmu *MMU) {

}

func opsORrd(z80 *Z80, mmu *MMU) {

}

func opsORre(z80 *Z80, mmu *MMU) {

}

func opsORrh(z80 *Z80, mmu *MMU) {

}

func opsORrl(z80 *Z80, mmu *MMU) {

}

func opsORHL(z80 *Z80, mmu *MMU) {

}

func opsORra(z80 *Z80, mmu *MMU) {

}

func opsCPrb(z80 *Z80, mmu *MMU) {

}

func opsCPrc(z80 *Z80, mmu *MMU) {

}

func opsCPrd(z80 *Z80, mmu *MMU) {

}

func opsCPre(z80 *Z80, mmu *MMU) {

}

func opsCPrh(z80 *Z80, mmu *MMU) {

}

func opsCPrl(z80 *Z80, mmu *MMU) {

}

func opsCPHL(z80 *Z80, mmu *MMU) {

}

func opsCPra(z80 *Z80, mmu *MMU) {

}

// C0
func opsRETNZ(z80 *Z80, mmu *MMU) {

}

func opsPOPBC(z80 *Z80, mmu *MMU) {

}

func opsJPNZnn(z80 *Z80, mmu *MMU) {

}

func opsJPnn(z80 *Z80, mmu *MMU) {

}

func opsCALLNZnn(z80 *Z80, mmu *MMU) {

}

func opsPUSHBC(z80 *Z80, mmu *MMU) {

}

func opsADDn(z80 *Z80, mmu *MMU) {

}

func opsRST00(z80 *Z80, mmu *MMU) {

}

func opsRETZ(z80 *Z80, mmu *MMU) {

}

func opsRET(z80 *Z80, mmu *MMU) {

}

func opsJPZnn(z80 *Z80, mmu *MMU) {

}

func opsMAPcb(z80 *Z80, mmu *MMU) {

}

func opsCALLZnn(z80 *Z80, mmu *MMU) {

}

func opsCALLnn(z80 *Z80, mmu *MMU) {

}

func opsADCn(z80 *Z80, mmu *MMU) {

}

func opsRST08(z80 *Z80, mmu *MMU) {

}

func opsRETNC(z80 *Z80, mmu *MMU) {

}

func opsPOPDE(z80 *Z80, mmu *MMU) {

}

func opsJPNCnn(z80 *Z80, mmu *MMU) {

}

func opsCALLNCnn(z80 *Z80, mmu *MMU) {

}

func opsPUSHDE(z80 *Z80, mmu *MMU) {

}

func opsSUBn(z80 *Z80, mmu *MMU) {

}

func opsRST10(z80 *Z80, mmu *MMU) {

}

func opsRETC(z80 *Z80, mmu *MMU) {

}

func opsRETI(z80 *Z80, mmu *MMU) {

}

func opsJPCnn(z80 *Z80, mmu *MMU) {

}

func opsCALLCnn(z80 *Z80, mmu *MMU) {

}

func opsSBCn(z80 *Z80, mmu *MMU) {

}

func opsRST18(z80 *Z80, mmu *MMU) {

}

// E0
func opsLDIOnA(z80 *Z80, mmu *MMU) {

}

func opsPOPHL(z80 *Z80, mmu *MMU) {

}

func opsLDIOCA(z80 *Z80, mmu *MMU) {

}

func opsPUSHHL(z80 *Z80, mmu *MMU) {

}

func opsANDn(z80 *Z80, mmu *MMU) {

}

func opsRST20(z80 *Z80, mmu *MMU) {

}

func opsADDSPn(z80 *Z80, mmu *MMU) {

}

func opsJPHL(z80 *Z80, mmu *MMU) {

}

func opsLDmmA(z80 *Z80, mmu *MMU) {

}

func opsORn(z80 *Z80, mmu *MMU) {

}

func opsRST28(z80 *Z80, mmu *MMU) {

}

func opsLDAIOn(z80 *Z80, mmu *MMU) {

}

func opsPOPAF(z80 *Z80, mmu *MMU) {

}

func opsLDAIOC(z80 *Z80, mmu *MMU) {

}

func opsDI(z80 *Z80, mmu *MMU) {

}

func opsPUSHAF(z80 *Z80, mmu *MMU) {

}

func opsXORn(z80 *Z80, mmu *MMU) {

}

func opsRST30(z80 *Z80, mmu *MMU) {

}

func opsLDHLSPn(z80 *Z80, mmu *MMU) {

}

func opsLDAmm(z80 *Z80, mmu *MMU) {

}

func opsEI(z80 *Z80, mmu *MMU) {

}

func opsCPn(z80 *Z80, mmu *MMU) {

}

func opsRST38(z80 *Z80, mmu *MMU) {

}
