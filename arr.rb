

pre = "spoints :=[]*points{"
suf = "}"
arr = (0..61).to_a
con = arr.map{|i| "s#{i}"}.join(",")


puts pre + con + suf
