# Wallet System Design
# ===
# Imagine that you are running a project that helps customers manage
# their money inside your company.
# ### Requirements
# - Customers can transfer their money to others
# - Customers are able to check their own wallet balance
# - Customers are able to check their transfers history, 
#  where history contains name of sender and receiver
# ### Question:
# - How would you design from Database layer to API layer

1. database
2. api

--
customer -> id, name, walletbalance
trnmoneytransfer -> id, date, code, amount, userfrom, userto

RESTRICT

Replace & Delete
id joshua
ke id dwi

{
    UserFrom: 'Juan'
    UserTo: 'Joshua'
    Amount: 10000
    Date: '2022-09-22'
}

function sendMoney(Request $request) {
    db::beginTransaction();
    try {
        $validate->UserFrom;
        $validate->UserTo;
        $validate->Amount;
        $validate->Date;

        $data = new MoneyTransfer();
        $customer = Customer::where('name', $request->UserFrom)->first();
        $transferto = Customer::where('name', $request->UserTo)->first();

        if ($customer->WalletBalance - $data->Amount < 0) throw exception('ga ada duit mas');
        $data->Amount = $request->Amount;
        $data->Date = $request->Date;
        $data->UserFrom = $customer->id;
        $data->UserTo = $transferto->id;
        $data->save();
        
        $customer->WalletBalance = $customer->WalletBalance - $data->Amount;
        $customer->save();

        $transferto->WalletBalance = $transferto->WalletBalance + $data->Amount;
        $transferto->save();

        DB::commit();

        return $data;
    } catch {
        db::rollback();
        sendtokanban(err)
    }
    
}

1. button check balance (GET)
'/checkbalance?id=joshua'

function checkWalletBalance(Request $request) {
    try {
        $customer = Customer::where('name', $request->id)->first();
        if (!$customer) throwcatch("customer doesnt exist");

        return $customer->WalletBalance;
    } catch {
        catch
    }
}

1. button check transfer history (GET)
'/checktransferhistory?id=joshua'

function checkTransferHistory(Request $request) {
    try {
        $query = "SELECT * FROM (
                    SELECT p.id, p.date, p.amount, cfrom.name UserFrom, cto.name UserTo FROM trnmoneytransfer p
                    LEFT JOIN customer cfrom ON cfrom.id = p.UserFrom
                    LEFT JOIN customer cto ON cto.id = p.UserTo
                    WHERE p.UserFrom = ?

                    UNION ALL
                    SELECT p.id, p.date, p.amount, cfrom.name UserFrom, cto.name UserTo FROM trnmoneytransfer p
                    LEFT JOIN customer cfrom ON cfrom.id = p.UserFrom
                    LEFT JOIN customer cto ON cto.id = p.UserTo
                    WHERE p.UserTo = ?
                ) p
                ORDER BY p.Date DESC
                ";
        $data = DB::SELECT($query, [$request->id, $request->id]);
        $result = [];
        foreach ($data as $row) {
            $result[] = [
                'id' => $row->id,
                'date' => $row->Date,
                'amount' => $row->Amount,
                'userfrom'=> $row->UserFrom,
                'userto'=> $row->UserTo
            ];
        }

        return $result;
    } catch {
        catch
    }
}

walletamount dari range 100rb - 200rb

$query = "SELECT * FROM customer 
    WHERE IFNULL(WalletBalance,0) >= 100000 
    AND IFNULL(WalletBalance,0) <= 200000";

apabila user ada 500 ribu.

Yang dicari adalah B

L > B
L < B

A - L
F

B
--
Field Gender
segmentasi customer yang pria, balance 500rb.

insert to customer (id, gender(indexed), balance(indexed))
1. insert customer
2. insert gender (L,P)
3. insert balance (500.000)

L
L
L
L

100
200
300
400
500
600

oid, name, username, password
oid -> uuid()
name -> request->name
username -> request->username
password -> hash
---

1. function save() {
}

2. function moneyTransfer() {
    tulis logic disini.

    ->save($request);
}



